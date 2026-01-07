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

import com.anb.admin.domain.Aptsubmaster;
import com.anb.admin.domain.AptsubmasterRepository;

@Service
public class AptsubmasterService {

    @Autowired
    AptsubmasterRepository repository;

    @Transactional
    public Aptsubmaster insert(Aptsubmaster item) {
        return repository.save(item);
    }

    @Transactional
    public Aptsubmaster update(Aptsubmaster item) {
        return repository.save(item);
    }

    @Transactional
    public void delete(Aptsubmaster item) {
        repository.delete(item);
    }

    public Optional<Aptsubmaster> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Aptsubmaster> findAll(int page, int size) {
        Pageable pageable = PageRequest.of(page, size, Sort.by("id").descending());
        return repository.findAll(pageable);
    }

    public List<Aptsubmaster> findByApt(Long apt) {
        return repository.findByApt(apt);
    }

    @Transactional
    public void deleteByApt(Long apt) {
        repository.deleteByApt(apt);
    }
}
