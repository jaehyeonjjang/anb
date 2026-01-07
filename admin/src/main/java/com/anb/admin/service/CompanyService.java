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
    
import com.anb.admin.domain.Company;
import com.anb.admin.domain.CompanyRepository;
import com.anb.admin.domain.CompanySpecs;
import com.anb.admin.domain.CompanySpecs.SearchKey;
import com.anb.admin.domain.Statuscategory;
import com.anb.admin.domain.StatuscategoryRepository;
import com.anb.admin.domain.Status;
import com.anb.admin.domain.StatusRepository;

@Service
public class CompanyService {

    @Autowired
    CompanyRepository repository;

    @Autowired
    StatuscategoryRepository statuscategoryRepository;

    @Autowired
    StatusRepository statusRepository;

    @Transactional
    public Company insert(Company item) {
        Optional<Company> opt = repository.findByName(item.getName());
        if (opt.isPresent()) {
            return null;
        }

        Company company = repository.save(item);
        Long id = company.getId();

        List<Statuscategory> statuscategorys = statuscategoryRepository.findByCompany(0L);
        for (Statuscategory statuscategory : statuscategorys) {
            Statuscategory newItem = new Statuscategory();
            newItem.setName(statuscategory.getName());
            newItem.setType(statuscategory.getType());
            newItem.setOrder(statuscategory.getOrder());
            newItem.setCompany(id);
            Statuscategory saveItem = statuscategoryRepository.save(newItem);

            List<Status> statuss = statusRepository.findByCompanyAndStatuscategory(0L, statuscategory.getId());
            for (Status status : statuss) {
                Status newStatus = new Status();
                newStatus.setName(status.getName());
                newStatus.setStatuscategory(saveItem.getId());
                newStatus.setType(status.getType());
                newStatus.setContent(status.getContent());
                newStatus.setEtc(status.getEtc());
                newStatus.setOrder(status.getOrder());
                newStatus.setCompany(id);
                statusRepository.save(newStatus);
            }
        }

        List<Status> statuss = statusRepository.findByCompanyAndStatuscategory(0L, 0L);
        for (Status status : statuss) {
            Status newStatus = new Status();
            newStatus.setName(status.getName());
            newStatus.setStatuscategory(0L);
            newStatus.setType(status.getType());
            newStatus.setContent(status.getContent());
            newStatus.setEtc(status.getEtc());
            newStatus.setOrder(status.getOrder());
            newStatus.setCompany(id);
            statusRepository.save(newStatus);
        }
        

        return company;
    }

    @Transactional
    public Company update(Company item) {
        Optional<Company> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        opt = repository.findByName(item.getName());

        if (opt.isPresent()) {
            if (item.getId() != opt.get().getId()) {
                return null;
            }
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(Long id) {
        Optional<Company> opt = repository.findById(id);

        if (opt.isPresent()) {
            repository.delete(opt.get());
        }
    }

    public Optional<Company> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Company> findAll(Map<SearchKey, Object> searchKeys, int page, int size) {
        Pageable pageableWithSort = PageRequest.of(page, size, Sort.by("id").descending());

        return searchKeys.isEmpty()
            ? repository.findAll(pageableWithSort)
            : repository.findAll(CompanySpecs.searchWith(searchKeys), pageableWithSort);
    }

    public List<Company> findAll(Map<SearchKey, Object> searchKeys) {
        return repository.findAll(CompanySpecs.searchWith(searchKeys));
    }

    public List<Company> findByStatus(int status) {
        return repository.findByStatus(status);
    }
}
