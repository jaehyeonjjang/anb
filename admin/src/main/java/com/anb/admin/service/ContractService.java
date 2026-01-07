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
    
import com.anb.admin.domain.Contract;
import com.anb.admin.domain.ContractRepository;
import com.anb.admin.domain.ContractSpecs;
import com.anb.admin.domain.ContractSpecs.SearchKey;
import com.anb.admin.domain.Company;
import com.anb.admin.domain.CompanyRepository;

@Service
public class ContractService {

    @Autowired
    ContractRepository repository;

    @Autowired
    CompanyRepository companyRepository;

    @Transactional
    public Contract insert(Contract item) {
        return repository.save(item);
    }

    @Transactional
    public Contract update(Contract item) {
        if (item.getStatus() == 2) {
            Optional<Company> opt = companyRepository.findById(item.getCompany());

            if (opt.isPresent()) {
                Company company = opt.get();

                company.setContractstartdate(item.getContractstartdate());
                company.setContractenddate(item.getContractenddate());

                companyRepository.save(company);
            }
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(Long id) {
        Optional<Contract> opt = repository.findById(id);

        if (opt.isPresent()) {
            repository.delete(opt.get());
        }
    }

    public Optional<Contract> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Contract> findAll(Map<SearchKey, Object> searchKeys, int page, int size) {
        Pageable pageableWithSort = PageRequest.of(page, size, Sort.by("id").descending());

        return searchKeys.isEmpty()
            ? repository.findAll(pageableWithSort)
            : repository.findAll(ContractSpecs.searchWith(searchKeys), pageableWithSort);
    }

    public List<Contract> findByStatus(int status) {
        return repository.findByStatusOrderById(status);
    }

    public List<Contract> findByCompanyAndStatus(Long company, int status) {
        return repository.findByCompanyAndStatusOrderById(company, status);
    }

    public List<Contract> findByCompany(Long company) {
        return repository.findByCompanyOrderById(company);
    }
}
