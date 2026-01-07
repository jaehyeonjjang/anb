package com.anb.admin.domain;

import java.util.Optional;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface SyncRepository extends PagingAndSortingRepository<Sync, Long>, JpaSpecificationExecutor<Sync> {
}
