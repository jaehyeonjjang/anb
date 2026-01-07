package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Sort;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface StatusRepository extends PagingAndSortingRepository<Status, Long>, JpaSpecificationExecutor<Status> {
    List<Status> findByCompany(Long company);
    List<Status> findByCompanyAndStatuscategory(Long company, Long statuscategory);
    Optional<Status> findByCompanyAndTypeAndStatuscategoryAndName(Long company, int type, Long statuscategory, String name);
}
