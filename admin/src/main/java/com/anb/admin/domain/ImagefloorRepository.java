package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface ImagefloorRepository extends PagingAndSortingRepository<Imagefloor, Long>, JpaSpecificationExecutor<Imagefloor> {
    Optional<Imagefloor> findByImageAndNameAndImagename(Long image, String name, String imagename);
    Optional<Imagefloor> findByTarget(Long target);
    List<Imagefloor> findByImage(Long image);
}
