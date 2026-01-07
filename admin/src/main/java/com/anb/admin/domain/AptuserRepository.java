package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface AptuserRepository extends PagingAndSortingRepository<Aptuser, Long>, JpaSpecificationExecutor<Aptuser> {
    List<Aptuser> findByApt(Long apt);
    List<Aptuser> findByUser(Long user);
    Page<Aptuser> findByAptAndLevel(Long apt, int level, Pageable pageable);
    Optional<Aptuser> findFirstByAptAndUser(Long apt, Long user);
    void deleteByApt(Long apt);
}
