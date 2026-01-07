package com.anb.admin.service;

import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
    
import com.anb.admin.domain.Sync;
import com.anb.admin.domain.SyncRepository;

@Service
public class SyncService {

    @Autowired
    SyncRepository repository;

    public Optional<Sync> findById(Long id) {
        return repository.findById(id);
    }
    
    public Page<Sync> findAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    @Transactional
    public Sync insert(Sync item) {
        return repository.save(item);
    }

    @Transactional
    public Sync update(Sync item) {
        return repository.save(item);
    }

    @Transactional
    public void delete(Sync item) {
        repository.delete(item);
    }
}
