package com.anb.admin.api;

import java.util.Optional;
import java.util.List;

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

import com.anb.admin.domain.Work;
import com.anb.admin.service.WorkService;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestController
@RequestMapping("/api/work")
public class WorkController {

    @Autowired
    WorkService service;

    @GetMapping("/{id}")
    public Work findById(@PathVariable Long id) {
        Optional<Work> opt = service.findById(id);
        return opt.isPresent() ? opt.get() : new Work();
    }

    @GetMapping("")
    public Page<Work> findAll(Pageable pageable) {
        return service.findAll(pageable);
    }

    @PostMapping("")
    public Work insert(@RequestBody Work item) {
        return service.insert(item);
    }

    @PutMapping("{id}")
    public Work update(@PathVariable Long id, @RequestBody Work item) {
        Optional<Work> opt = service.findById(id);

        if (opt.isPresent()) {
            return service.update(item);
        }

        return null;
    }

    @DeleteMapping("{id}")
    public void insert(@PathVariable Long id) {
        Optional<Work> opt = service.findById(id);

        if (opt.isPresent()) {
            service.delete(opt.get());
        }
    }
}
