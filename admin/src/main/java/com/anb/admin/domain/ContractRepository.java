package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface ContractRepository extends PagingAndSortingRepository<Contract, Long>, JpaSpecificationExecutor<Contract> {
    List<Contract> findByStatusOrderById(int status);
    List<Contract> findByCompanyOrderById(Long company);
    List<Contract> findByCompanyAndStatusOrderById(Long company, int status);
}
