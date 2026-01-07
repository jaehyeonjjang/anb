package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface AptRepository extends PagingAndSortingRepository<Apt, Long>, JpaSpecificationExecutor<Apt> {
    List<Apt> findByCompany(Long company);
    List<Apt> findByCompanyAndStatus(Long company, int tatus);
    List<Apt> findByAptgroup(Long aptgroup);
    List<Apt> findByCompanyAndReport(Long company, int report);
}
