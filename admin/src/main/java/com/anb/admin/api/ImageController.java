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

import com.anb.admin.domain.Image;
import com.anb.admin.service.ImageService;

@Slf4j
@RestController
@RequestMapping("/api/image")
public class ImageController {

    @Autowired
    ImageService service;

    @PostMapping("")
    public ResponseEntity<? extends BasicResponse> insert(@RequestBody Image item) {
        Image result = service.insert(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Image>(result));
    }

    @PutMapping("{id}")
    public ResponseEntity<? extends BasicResponse> update(@PathVariable Long id, @RequestBody Image item) {
        Image result = service.update(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Image>(result));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<? extends BasicResponse> delete(@PathVariable Long id) {
        Optional<Image> opt = service.findById(id);

        if (opt.isPresent()) {
            service.delete(opt.get());
        }

        return ResponseEntity.noContent().build();
    }

    @GetMapping("{id}")
    public ResponseEntity<? extends BasicResponse> findById(@PathVariable Long id) {
        Optional<Image> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.ok().body(new ErrorResponse("데이터를 찾을수가 없습니다", HttpStatus.NOT_FOUND));
        }

        return ResponseEntity.ok().body(new CommonResponse<Image>(opt.get()));
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "apt") Long apt) {
        List<Image> result = service.findByApt(apt);
        return ResponseEntity.ok().body(new CommonResponse<List<Image>>(result));        
    }
}
