package com.anb.admin.service;

import java.util.Optional;
import java.util.List;
import java.util.Map;

import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Sort;
import org.springframework.data.domain.PageRequest;
import org.apache.commons.lang3.StringUtils;

import com.anb.admin.domain.User;
import com.anb.admin.domain.UserRepository;
import com.anb.admin.domain.UserSpecs;
import com.anb.admin.domain.UserSpecs.SearchKey;

@Service
public class UserService {

    @Autowired
    UserRepository repository;

    public String sha256(String msg) {
       try {
           MessageDigest md = MessageDigest.getInstance("SHA-256");
           md.update(msg.getBytes());

           byte[] bytes = md.digest();
           StringBuilder builder = new StringBuilder();
           for (byte b: bytes) {
               builder.append(String.format("%02x", b));
           }
           return builder.toString();
       } catch (NoSuchAlgorithmException e) {
       }

       return "";
    }

    @Transactional
    public User insert(User item) {
        Optional<User> opt = repository.findByLoginid(item.getLoginid());

        if (opt.isPresent()) {
            return null;
        }

        String passwd = sha256(item.getPasswd());
        item.setPasswd(passwd);

        return repository.save(item);
    }

    @Transactional
    public User update(User item) {
        Optional<User> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        User old = opt.get();

        opt = repository.findByLoginid(item.getLoginid());

        if (opt.isPresent()) {
            if (old.getId() != opt.get().getId()) {
                return null;
            }
        }

        String passwd = item.getPasswd();
        if (StringUtils.isEmpty(passwd)) {
            item.setPasswd(old.getPasswd());
        } else {
            item.setPasswd(sha256(passwd));
        }

        return repository.save(item);
    }

    @Transactional
    public void delete(User item) {
        repository.delete(item);
    }

    public Optional<User> findById(Long id) {
        return repository.findById(id);
    }

    public Page<User> findAll(Map<SearchKey, Object> searchKeys, String order, int page, int size) {
        Sort sort = null;
        boolean desc = false;

        if (StringUtils.isEmpty(order)) {
            order = "id";
        } else {
            if (StringUtils.right(order, 4).equals("Desc")) {
                order = StringUtils.left(order, order.length() - 4);
                desc = true;
            }
        }

        sort = Sort.by(order);

        if (desc) {
            sort = sort.descending();
        }

        Pageable pageable = PageRequest.of(page, size, sort);

        return searchKeys.isEmpty()
            ? repository.findAll(pageable)
            : repository.findAll(UserSpecs.searchWith(searchKeys), pageable);
    }

    public List<User> findByCompany(Long company) {
        return repository.findByCompany(company);
    }

    public List<User> findByCompanyAndStatus(Long company, int status) {
        return repository.findByCompanyAndStatus(company, status);
    }

    public Optional<User> findByLoginid(String loginid) {
        return repository.findByLoginid(loginid);
    }
}
