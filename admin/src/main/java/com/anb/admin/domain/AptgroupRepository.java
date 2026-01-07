package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface AptgroupRepository extends PagingAndSortingRepository<Aptgroup, Long>, JpaSpecificationExecutor<Aptgroup> {
    List<Aptgroup> findByCompanyAndStatus(Long company, int status);
    List<Aptgroup> findByCompany(Long company);
    Optional<Aptgroup> findByCompanyAndName(Long company, String name);
}
