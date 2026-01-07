package com.anb.admin.api;

import java.util.Optional;
import java.util.List;

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

import com.anb.admin.domain.Aptsubmaster;
import com.anb.admin.service.AptsubmasterService;

@Slf4j
@RestController
@RequestMapping("/api/aptsubmaster")
public class AptsubmasterController {

    @Autowired
    AptsubmasterService service;

    @PostMapping("")
    public Aptsubmaster insert(@RequestBody Aptsubmaster item) {
        return service.insert(item);
    }

    @PutMapping("{id}")
    public Aptsubmaster update(@PathVariable Long id, @RequestBody Aptsubmaster item) {
        Optional<Aptsubmaster> opt = service.findById(id);

        if (opt.isPresent()) {
            return service.update(item);
        }

        return null;
    }

    @DeleteMapping("{id}")
    public void delete(@PathVariable Long id) {
        Optional<Aptsubmaster> opt = service.findById(id);

        if (opt.isPresent()) {
            service.delete(opt.get());
        }
    }

    @GetMapping("/{id}")
    public Aptsubmaster findById(@PathVariable Long id) {
        Optional<Aptsubmaster> opt = service.findById(id);
        return opt.isPresent() ? opt.get() : new Aptsubmaster();
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "page", defaultValue = "0", required = false) int page,
                                                           @RequestParam(value = "size", defaultValue = "0", required = false) int size,
                                                           @RequestParam(value = "apt", defaultValue = "0", required = false) Long apt,
                                                           @RequestParam(value = "submaster", defaultValue = "0", required = false) Long submaster) {
        if (apt > 0) {
            List<Aptsubmaster> result = service.findByApt(apt);
            return ResponseEntity.ok().body(new CommonResponse<List<Aptsubmaster>>(result));
        } else {
            Page<Aptsubmaster> result = service.findAll(page, size);
            return ResponseEntity.ok().body(new CommonResponse<Page<Aptsubmaster>>(result));
        }
    }

    @DeleteMapping("apt/{id}")
    public void deleteByApt(@PathVariable Long id) {
        service.deleteByApt(id);
    }
}
