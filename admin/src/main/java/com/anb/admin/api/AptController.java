package com.anb.admin.api;

import java.util.Optional;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

import lombok.extern.slf4j.Slf4j;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;

import org.apache.commons.lang3.StringUtils;

import com.anb.admin.domain.Apt;
import com.anb.admin.domain.AptSpecs.SearchKey;
import com.anb.admin.service.AptService;

@Slf4j
@RestController
@RequestMapping("/api/apt")
public class AptController {

    @Autowired
    AptService service;

    @PostMapping("")
    public ResponseEntity<? extends BasicResponse> insert(@RequestBody Apt item) {
        Apt result = service.insert(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Apt>(result));
    }

    @PutMapping("{id}")
    public ResponseEntity<? extends BasicResponse> update(@PathVariable Long id, @RequestBody Apt item) {
        Apt result = service.update(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Apt>(result));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<? extends BasicResponse> delete(@PathVariable Long id) {
        Optional<Apt> opt = service.findById(id);

        if (opt.isPresent()) {
            service.delete(opt.get());
        }

        return ResponseEntity.noContent().build();
    }

    @GetMapping("{id}")
    public ResponseEntity<? extends BasicResponse> findById(@PathVariable Long id) {
        Optional<Apt> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.ok().body(new ErrorResponse("데이터를 찾을수가 없습니다", HttpStatus.NOT_FOUND));
        }

        return ResponseEntity.ok().body(new CommonResponse<Apt>(opt.get()));
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "page", defaultValue = "0", required = false) int page,
                                                           @RequestParam(value = "size", defaultValue = "0", required = false) int size,
                                                           @RequestParam(value = "orderby", required = false) String orderby,
                                                           @RequestParam(value = "status", defaultValue = "0", required = false) int status,
                                                           @RequestParam(value = "aptgroup", defaultValue = "0", required = false) Long aptgroup,
                                                           @RequestParam(value = "company", defaultValue = "0", required = false) Long company,
                                                           @RequestParam(value = "report", defaultValue = "0", required = false) int report,
                                                           @RequestParam(value = "name", required = false) String name) {
        Map<SearchKey, Object> searchKeys = new HashMap<>();
        if (status > 0) {
            searchKeys.put(SearchKey.valueOf("STATUS"), status);
        }

        if (aptgroup > 0) {
            searchKeys.put(SearchKey.valueOf("APTGROUP"), aptgroup);
        }

        if (company > 0) {
            searchKeys.put(SearchKey.valueOf("COMPANY"), company);
        }

        if (!StringUtils.isEmpty(name)) {
            searchKeys.put(SearchKey.valueOf("SEARCH"), name);
        }

        if (size == 0) {
            if (aptgroup > 0) {
                List<Apt> result = service.findByAptgroup(aptgroup);
                return ResponseEntity.ok().body(new CommonResponse<List<Apt>>(result));
            } else {
                if (report > 0) {
                    List<Apt> result = service.findByCompanyAndReport(company, report);
                    return ResponseEntity.ok().body(new CommonResponse<List<Apt>>(result));
                } else {
                    List<Apt> result = service.findByCompany(company);
                    return ResponseEntity.ok().body(new CommonResponse<List<Apt>>(result));
                }
            }
        } else {
            Page<Apt> result = service.findAll(searchKeys, orderby, page, size);
            return ResponseEntity.ok().body(new CommonResponse<Page<Apt>>(result));
        }
    }

    @PostMapping("{id}/copy")
    public ResponseEntity<? extends BasicResponse> copy(@PathVariable Long id) {
        Optional<Apt> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.ok().body(new ErrorResponse("데이터를 찾을수가 없습니다", HttpStatus.NOT_FOUND));
        }

        Apt result = service.copy(id);

        return ResponseEntity.ok().body(new CommonResponse<Apt>(result));
    }
}
