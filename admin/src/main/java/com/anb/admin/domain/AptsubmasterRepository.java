package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface AptsubmasterRepository extends PagingAndSortingRepository<Aptsubmaster, Long>, JpaSpecificationExecutor<Aptsubmaster> {
    List<Aptsubmaster> findByApt(Long apt);
    List<Aptsubmaster> findByUser(Long user);
    Page<Aptsubmaster> findByAptAndLevel(Long apt, int level, Pageable pageable);
    Optional<Aptsubmaster> findFirstByAptAndUser(Long apt, Long user);
    void deleteByApt(Long apt);
}
