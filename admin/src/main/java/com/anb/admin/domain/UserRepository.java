package com.anb.admin.domain;

import java.util.Optional;
import java.util.List;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.repository.PagingAndSortingRepository;

public interface UserRepository extends PagingAndSortingRepository<User, Long>, JpaSpecificationExecutor<User> {
    List<User> findByCompany(Long company);
    List<User> findByCompanyAndStatus(Long company, int status);
    Optional<User> findById(Long id);    
    Optional<User> findByLoginid(String loginid);
}
