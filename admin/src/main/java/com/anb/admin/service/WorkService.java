package com.anb.admin.service;

import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
    
import com.anb.admin.domain.Work;
import com.anb.admin.domain.WorkRepository;

@Service
public class WorkService {

    @Autowired
    WorkRepository repository;

    public Optional<Work> findById(Long id) {
        return repository.findById(id);
    }
    
    public Page<Work> findAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    @Transactional
    public Work insert(Work item) {
        return repository.save(item);
    }

    @Transactional
    public Work update(Work item) {
        return repository.save(item);
    }

    @Transactional
    public void delete(Work item) {
        repository.delete(item);
    }
}
