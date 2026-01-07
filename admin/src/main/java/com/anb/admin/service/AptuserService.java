package com.anb.admin.service;

import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Sort;
import org.springframework.data.domain.PageRequest;
import org.apache.commons.lang3.StringUtils;

import com.anb.admin.domain.Aptuser;
import com.anb.admin.domain.AptuserRepository;

@Service
public class AptuserService {

    @Autowired
    AptuserRepository repository;

    @Transactional
    public Aptuser insert(Aptuser item) {
        return repository.save(item);
    }

    @Transactional
    public Aptuser update(Aptuser item) {
        return repository.save(item);
    }

    @Transactional
    public void delete(Aptuser item) {
        repository.delete(item);
    }

    public Optional<Aptuser> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Aptuser> findAll(int page, int size) {
        Pageable pageable = PageRequest.of(page, size, Sort.by("id").descending());
        return repository.findAll(pageable);
    }

    public List<Aptuser> findByApt(Long apt) {
        return repository.findByApt(apt);
    }

    @Transactional
    public void deleteByApt(Long apt) {
        repository.deleteByApt(apt);
    }
}
