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

import com.anb.admin.domain.User;
import com.anb.admin.domain.UserSpecs.SearchKey;
import com.anb.admin.service.UserService;

@Slf4j
@RestController
@RequestMapping("/api/user")
public class UserController {

    @Autowired
    UserService service;

    @PostMapping("")
    public ResponseEntity<? extends BasicResponse> insert(@RequestBody User item) {
        User result = service.insert(item);

        if (result == null) {
            return ResponseEntity.status(HttpStatus.CONFLICT).body(new ErrorResponse("이미 등록된 이메일입니다"));
        }

        return ResponseEntity.ok().body(new CommonResponse<User>(result));
    }

    @PutMapping("{id}")
    public ResponseEntity<? extends BasicResponse> update(@PathVariable Long id, @RequestBody User item) {
        Optional<User> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(new ErrorResponse("데이터를 찾을수가 없습니다"));
        }

        User result = service.update(item);
        return ResponseEntity.ok().body(new CommonResponse<User>(result));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<? extends BasicResponse> delete(@PathVariable Long id) {
        Optional<User> opt = service.findById(id);

        if (opt.isPresent()) {
            service.delete(opt.get());
        }

        return ResponseEntity.noContent().build();
    }

        @GetMapping("/{id}")
    public ResponseEntity<? extends BasicResponse> findById(@PathVariable Long id) {
        Optional<User> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(new ErrorResponse("데이터를 찾을수가 없습니다"));
        }

        User item = opt.get();
        item.setPasswd("");

        return ResponseEntity.ok().body(new CommonResponse<User>(item));
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "page", defaultValue = "0", required = false) int page,
                                                           @RequestParam(value = "size", defaultValue = "0", required = false) int size,
                                                           @RequestParam(value = "orderby", required = false) String orderby,
                                                           @RequestParam(value = "level", defaultValue = "0", required = false) int level,
                                                           @RequestParam(value = "status", defaultValue = "0", required = false) int status,
                                                           @RequestParam(value = "company", defaultValue = "0", required = false) Long company,
                                                           @RequestParam(value = "loginid", required = false) String loginid,
                                                           @RequestParam(value = "name", required = false) String name) {
        Map<SearchKey, Object> searchKeys = new HashMap<>();
        if (level > 0) {
            searchKeys.put(SearchKey.valueOf("LEVEL"), level);
        }

        if (status > 0) {
            searchKeys.put(SearchKey.valueOf("STATUS"), status);
        }

        if (company > 0) {
            searchKeys.put(SearchKey.valueOf("COMPANY"), company);
        }

        if (!StringUtils.isEmpty(loginid)) {
            searchKeys.put(SearchKey.valueOf("LOGINID"), loginid);
        }

        if (!StringUtils.isEmpty(name)) {
            searchKeys.put(SearchKey.valueOf("NAME"), name);
        }

        if (size == 0) {
            List<User> result = service.findByCompanyAndStatus(company, 1);
            return ResponseEntity.ok().body(new CommonResponse<List<User>>(result));
        } else {
            Page<User> result = service.findAll(searchKeys, orderby, page, size);
            return ResponseEntity.ok().body(new CommonResponse<Page<User>>(result));
        }
    }
}
