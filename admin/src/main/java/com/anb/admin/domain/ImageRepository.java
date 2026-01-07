package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface ImageRepository extends PagingAndSortingRepository<Image, Long>, JpaSpecificationExecutor<Image> {
    List<Image> findByAptOrderByOrderAscIdAsc(Long apt);
    List<Image> findByAptAndLevelLessThan(Long apt, int level);
    List<Image> findByAptAndParent(Long apt, Long parent);
}
