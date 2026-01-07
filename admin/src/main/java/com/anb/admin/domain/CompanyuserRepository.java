package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface CompanyuserRepository extends PagingAndSortingRepository<Companyuser, Long>, JpaSpecificationExecutor<Companyuser> {
    Page<Companyuser> findByCompany(Long company, Pageable pageable);
    Page<Companyuser> findByUser(Long user, Pageable pageable);
    Optional<Companyuser> findFirstByCompanyAndUser(Long company, Long user);
}
