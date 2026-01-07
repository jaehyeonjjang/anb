package com.anb.admin.api;

import java.util.Optional;
import java.util.List;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;

import com.anb.admin.domain.User;
import com.anb.admin.service.UserService;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestController
@RequestMapping("/api/login")
public class LoginController {

    @Autowired
    UserService service;

    @RequestMapping(value = "", method = RequestMethod.GET, produces = "application/json; charset=utf8")
    public ResponseEntity<? extends BasicResponse> login(@RequestParam(value = "loginid") String loginid, @RequestParam(value = "passwd") String passwd) {
        Optional<User> opt = service.findByLoginid(loginid);
        if (!opt.isPresent()) {
            log.info("not found");
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(new ErrorResponse("사용자를 찾을수가 없습니다"));
        }

        User user = opt.get();

        log.info("passwd = " + passwd);
        log.info("user.passwd = " + user.getPasswd());

        return ResponseEntity.ok().body(new CommonResponse<User>(user));

        //return user;

        /*
        if (user.getPasswd().equals(passwd)) {
            return user;
        } else {
            log.info("wrong passwd");
            return new User();
        }
        */
    }

    @RequestMapping(value = "logout", method = RequestMethod.GET)
    public ResponseEntity<? extends BasicResponse> logout() {
        return ResponseEntity.noContent().build();
    }
}
