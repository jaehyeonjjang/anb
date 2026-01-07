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
import org.apache.commons.lang3.StringUtils;

import com.anb.admin.domain.Aptgroup;
import com.anb.admin.domain.AptgroupRepository;
import com.anb.admin.domain.AptgroupSpecs;
import com.anb.admin.domain.AptgroupSpecs.SearchKey;
import com.anb.admin.domain.Apt;
import com.anb.admin.domain.AptRepository;

@Service
public class AptgroupService {

    @Autowired
    AptgroupRepository repository;

    @Autowired
    AptRepository aptRepository;

    @Transactional
    public Aptgroup insert(Aptgroup item) {
        Optional<Aptgroup> opt = repository.findByCompanyAndName(item.getCompany(), item.getName());

        if (opt.isPresent()) {
            return null;
        }

        return repository.save(item);
    }

    @Transactional
    public Aptgroup update(Aptgroup item) {
        Optional<Aptgroup> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        Aptgroup old = opt.get();

        opt = repository.findByCompanyAndName(item.getCompany(), item.getName());

        if (opt.isPresent()) {
            if (old.getId() != opt.get().getId()) {
                return null;
            }
        }

        List<Apt> apts = aptRepository.findByAptgroup(item.getId());

        for (Apt apt : apts) {
            String search = item.getName() + " " + apt.getName();
            apt.setSearch(search);
            aptRepository.save(apt);
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(Aptgroup item) {
        repository.delete(item);
    }

    public Optional<Aptgroup> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Aptgroup> findAll(Map<SearchKey, Object> searchKeys, String order, int page, int size) {
        Sort sort = null;
        boolean desc = false;
        
        if (StringUtils.isEmpty(order)) {
            order = "id";            
        } else {
            if (StringUtils.right(order, 4).equals("Desc")) {
                order = StringUtils.left(order, order.length() - 4);
                desc = true;
            }
        }

        sort = Sort.by(order);

        if (desc) {
            sort = sort.descending();
        }
        
        Pageable pageable = PageRequest.of(page, size, sort);

        return searchKeys.isEmpty()
            ? repository.findAll(pageable)
            : repository.findAll(AptgroupSpecs.searchWith(searchKeys), pageable);
    }

    public List<Aptgroup> findByCompany(Long company) {
        return repository.findByCompany(company);
    }

    public List<Aptgroup> findByCompanyAndStatus(Long company, int status) {
        return repository.findByCompanyAndStatus(company, status);
    }
}
