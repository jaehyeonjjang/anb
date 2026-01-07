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
    
import com.anb.admin.domain.Report;
import com.anb.admin.domain.ReportRepository;
import com.anb.admin.domain.ReportSpecs;
import com.anb.admin.domain.ReportSpecs.SearchKey;
import com.anb.admin.domain.Apt;
import com.anb.admin.domain.AptRepository;

@Service
public class ReportService {

    @Autowired
    ReportRepository repository;

    @Autowired
    AptRepository aptRepository;

    @Transactional
    public Report insert(Report item) {
        Optional<Apt> opt = aptRepository.findById(item.getApt());

        if (!opt.isPresent()) {
            return null;
        }

        Apt apt = opt.get();

        if (apt.getReport() == 2 ||  apt.getReport() == 3) {
            return item;
        }

        apt.setReport(2);
        aptRepository.save(apt);

        return repository.save(item);
    }

    @Transactional
    public Report update(Report item) {
        Optional<Report> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(Long id) {
        Optional<Report> opt = repository.findById(id);

        if (opt.isPresent()) {
            repository.delete(opt.get());
        }
    }

    public Optional<Report> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Report> findAll(Map<SearchKey, Object> searchKeys, int page, int size) {
        Pageable pageableWithSort = PageRequest.of(page, size, Sort.by("id").descending());

        return searchKeys.isEmpty()
            ? repository.findAll(pageableWithSort)
            : repository.findAll(ReportSpecs.searchWith(searchKeys), pageableWithSort);
    }

    public List<Report> findAll(Map<SearchKey, Object> searchKeys) {
        return repository.findAll(ReportSpecs.searchWith(searchKeys));
    }

    public List<Report> findByStatus(int status) {
        return repository.findByStatus(status);
    }
}
