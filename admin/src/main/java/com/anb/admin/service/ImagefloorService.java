package com.anb.admin.service;

import java.util.Optional;
import java.util.List;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
    
import com.anb.admin.domain.Imagefloor;
import com.anb.admin.domain.ImagefloorRepository;

@Service
public class ImagefloorService {

    @Autowired
    ImagefloorRepository repository;

    public Optional<Imagefloor> findById(Long id) {
        return repository.findById(id);
    }
    
    public Page<Imagefloor> findAll(Pageable pageable) {
        return repository.findAll(pageable);
    }

    @Transactional
    public Imagefloor insert(Imagefloor item) {
        return repository.save(item);
    }

    @Transactional
    public Imagefloor update(Imagefloor item) {
        return repository.save(item);
    }

    @Transactional
    public void delete(Imagefloor item) {
        repository.delete(item);
    }

    public List<Imagefloor> findByImage(Long image) {
        return repository.findByImage(image);
    }
}
