package com.anb.admin.service;

import java.util.Optional;
import java.util.List;
import java.util.Map;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Sort;
import org.springframework.data.domain.PageRequest;
    
import com.anb.admin.domain.Statuscategory;
import com.anb.admin.domain.StatuscategoryRepository;
import com.anb.admin.domain.StatuscategorySpecs;
import com.anb.admin.domain.StatuscategorySpecs.SearchKey;

@Service
public class StatuscategoryService {

    @Autowired
    StatuscategoryRepository repository;

    @Transactional
    public Statuscategory insert(Statuscategory item) {
        Optional<Statuscategory> opt = repository.findByTypeAndName(item.getType(), item.getName());
        if (opt.isPresent()) {
            return null;
        }

        return repository.save(item);
    }

    @Transactional
    public Statuscategory update(Statuscategory item) {
        Optional<Statuscategory> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        opt = repository.findByTypeAndName(item.getType(), item.getName());

        if (opt.isPresent()) {
            if (item.getId() != opt.get().getId()) {
                return null;
            }
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(Long id) {
        Optional<Statuscategory> opt = repository.findById(id);

        if (opt.isPresent()) {
            repository.delete(opt.get());
        }
    }

    public Optional<Statuscategory> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Statuscategory> findAll(Map<SearchKey, Object> searchKeys, int page, int size) {
        Pageable pageable = PageRequest.of(page, size, Sort.by("id").descending());

        return searchKeys.isEmpty()
            ? repository.findAll(pageable)
            : repository.findAll(StatuscategorySpecs.searchWith(searchKeys), pageable);
    }

    public List<Statuscategory> findAll(Map<SearchKey, Object> searchKeys) {
        return repository.findAll(StatuscategorySpecs.searchWith(searchKeys), Sort.by("order").and(Sort.by("id")));
    }
}
