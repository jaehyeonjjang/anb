package com.anb.admin.service;

import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
    
import com.anb.admin.domain.Data;
import com.anb.admin.domain.DataRepository;

@Service
public class DataService {

    @Autowired
    DataRepository repository;

    public Optional<Data> findById(Long id) {
        return repository.findById(id);
    }
    
    public Page<Data> findAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    public List<Data> findByImage(Long image) {
        return repository.findByImage(image);
    }    

    @Transactional
    public Data insert(Data item) {
        return repository.save(item);
    }

    @Transactional
    public Data update(Data item) {
        return repository.save(item);
    }

    @Transactional
    public void delete(Data item) {
        repository.delete(item);
    }
}
