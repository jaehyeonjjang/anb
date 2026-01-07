package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface DataRepository extends PagingAndSortingRepository<Data, Long>, JpaSpecificationExecutor<Data> {
    List<Data> findByApt(Long apt, Pageable pageable);
    List<Data> findByImage(Long image);
    List<Data> findByImageAndNameAndImagename(Long image, String name, String imagename);
}
