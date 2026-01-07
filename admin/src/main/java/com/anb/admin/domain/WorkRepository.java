package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface WorkRepository extends PagingAndSortingRepository<Work, Long>, JpaSpecificationExecutor<Work> {
    Page<Work> findByImage(Long image, Pageable pageable);
}
