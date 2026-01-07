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

import com.anb.admin.domain.Status;
import com.anb.admin.domain.StatusSpecs.SearchKey;
import com.anb.admin.service.StatusService;

@Slf4j
@RestController
@RequestMapping("/api/status")
public class StatusController {

    @Autowired
    StatusService service;

    @PostMapping("")
    public ResponseEntity<? extends BasicResponse> insert(@RequestBody Status item) {
        Status result = service.insert(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Status>(result));
    }

    @PutMapping("{id}")
    public ResponseEntity<? extends BasicResponse> update(@PathVariable Long id, @RequestBody Status item) {
        Status result = service.update(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Status>(result));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<? extends BasicResponse> delete(@PathVariable Long id) {
        service.delete(id);

        return ResponseEntity.noContent().build();
    }

    @GetMapping("{id}")
    public ResponseEntity<? extends BasicResponse> findById(@PathVariable Long id) {
        Optional<Status> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.ok().body(new ErrorResponse("데이터를 찾을수가 없습니다", HttpStatus.NOT_FOUND));
        }

        return ResponseEntity.ok().body(new CommonResponse<Status>(opt.get()));
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "page", defaultValue = "0", required = false) int page,
                                                           @RequestParam(value = "size", defaultValue = "0", required = false) int size,
                                                           @RequestParam(value = "type", defaultValue = "0", required = false) int type,
                                                           @RequestParam(value = "statuscategory", required = false) Long statuscategory,
                                                           @RequestParam(value = "company", defaultValue = "0", required = false) Long company,
                                                           @RequestParam(value = "name", required = false) String name) {
        Map<SearchKey, Object> searchKeys = new HashMap<>();
        if (type > 0) {
            searchKeys.put(SearchKey.valueOf("TYPE"), type);
        }

        if (statuscategory > 0) {
            searchKeys.put(SearchKey.valueOf("STATUSCATEGORY"), statuscategory);
        }

        if (!StringUtils.isEmpty(name)) {
            searchKeys.put(SearchKey.valueOf("NAME"), name);
        }

        searchKeys.put(SearchKey.valueOf("COMPANY"), company);

        if (size == 0) {
            List<Status> result = service.findAll(searchKeys);
            return ResponseEntity.ok().body(new CommonResponse<List<Status>>(result));
        } else {
            Page<Status> result = service.findAll(searchKeys, page, size);
            return ResponseEntity.ok().body(new CommonResponse<Page<Status>>(result));
        }
    }
}
