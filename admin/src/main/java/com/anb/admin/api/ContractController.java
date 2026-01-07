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

import com.anb.admin.domain.Contract;
import com.anb.admin.domain.ContractSpecs.SearchKey;
import com.anb.admin.service.ContractService;

@Slf4j
@RestController
@RequestMapping("/api/contract")
public class ContractController {

    @Autowired
    ContractService service;

    @PostMapping("")
    public ResponseEntity<? extends BasicResponse> insert(@RequestBody Contract item) {
        Contract result = service.insert(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Contract>(result));
    }

    @PutMapping("{id}")
    public ResponseEntity<? extends BasicResponse> update(@PathVariable Long id, @RequestBody Contract item) {
        Contract result = service.update(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Contract>(result));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<? extends BasicResponse> delete(@PathVariable Long id) {
        service.delete(id);

        return ResponseEntity.noContent().build();
    }

    @GetMapping("{id}")
    public ResponseEntity<? extends BasicResponse> findById(@PathVariable Long id) {
        Optional<Contract> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.ok().body(new ErrorResponse("데이터를 찾을수가 없습니다", HttpStatus.NOT_FOUND));
        }

        return ResponseEntity.ok().body(new CommonResponse<Contract>(opt.get()));
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "page", defaultValue = "0", required = false) int page,
                                                           @RequestParam(value = "size", defaultValue = "0", required = false) int size,
                                                           @RequestParam(value = "status", defaultValue = "0", required = false) int status,
                                                           @RequestParam(value = "company", defaultValue = "0", required = false) Long company) {
        Map<SearchKey, Object> searchKeys = new HashMap<>();
        if (status > 0) {
            searchKeys.put(SearchKey.valueOf("STATUS"), status);
        }

        if (company > 0) {
            searchKeys.put(SearchKey.valueOf("COMPANY"), company);
        }

        Page<Contract> result = service.findAll(searchKeys, page, size);
        return ResponseEntity.ok().body(new CommonResponse<Page<Contract>>(result));
    }

    @GetMapping("search")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "status", defaultValue = "0", required = false) int status,
                                                           @RequestParam(value = "company", defaultValue = "0", required = false) Long company) {

        List<Contract> result = null;

        if (status > 0) {
            if (company > 0) {
                result = service.findByCompanyAndStatus(company, status);
            } else {
                result = service.findByStatus(status);
            }
        } else {
            if (company > 0) {
                result = service.findByCompany(company);
            }
        }

        return ResponseEntity.ok().body(new CommonResponse<List<Contract>>(result));
    }
}
