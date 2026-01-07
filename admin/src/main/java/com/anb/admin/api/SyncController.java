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

import com.anb.admin.domain.Sync;
import com.anb.admin.service.SyncService;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestController
@RequestMapping("/api/sync")
public class SyncController {

    @Autowired
    SyncService service;

    @GetMapping("/{id}")
    public Sync findById(@PathVariable Long id) {
        Optional<Sync> opt = service.findById(id);
        return opt.isPresent() ? opt.get() : new Sync();
    }

    @GetMapping("")
    public Page<Sync> findAll(Pageable pageable) {
        return service.findAll(pageable);
    }

    @PostMapping("")
    public Sync insert(@RequestBody Sync item) {
        return service.insert(item);
    }

    @PutMapping("{id}")
    public Sync update(@PathVariable Long id, @RequestBody Sync item) {
        Optional<Sync> opt = service.findById(id);

        if (opt.isPresent()) {
            return service.update(item);
        }

        return null;
    }

    @DeleteMapping("{id}")
    public void insert(@PathVariable Long id) {
        Optional<Sync> opt = service.findById(id);

        if (opt.isPresent()) {
            service.delete(opt.get());
        }
    }
}
