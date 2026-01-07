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
    
import com.anb.admin.domain.Status;
import com.anb.admin.domain.StatusRepository;
import com.anb.admin.domain.StatusSpecs;
import com.anb.admin.domain.StatusSpecs.SearchKey;

@Service
public class StatusService {

    @Autowired
    StatusRepository repository;

    @Transactional
    public Status insert(Status item) {
        Optional<Status> opt = repository.findByCompanyAndTypeAndStatuscategoryAndName(item.getCompany(), item.getType(), item.getStatuscategory(), item.getName());
        if (opt.isPresent()) {
            return null;
        }

        return repository.save(item);
    }

    @Transactional
    public Status update(Status item) {
        Optional<Status> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        opt = repository.findByCompanyAndTypeAndStatuscategoryAndName(item.getCompany(), item.getType(), item.getStatuscategory(), item.getName());

        if (opt.isPresent()) {
            if (item.getId() != opt.get().getId()) {
                return null;
            }
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(Long id) {
        Optional<Status> opt = repository.findById(id);

        if (opt.isPresent()) {
            repository.delete(opt.get());
        }
    }

    public Optional<Status> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Status> findAll(Map<SearchKey, Object> searchKeys, int page, int size) {
        Pageable pageable = PageRequest.of(page, size, Sort.by("id").descending());

        return searchKeys.isEmpty()
            ? repository.findAll(pageable)
            : repository.findAll(StatusSpecs.searchWith(searchKeys), pageable);
    }

    public List<Status> findAll(Map<SearchKey, Object> searchKeys) {
        return repository.findAll(StatusSpecs.searchWith(searchKeys), Sort.by("order").and(Sort.by("id")));
    }
}
