package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="user_tb")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "u_id")
    private Long id;

    @Column(name = "u_loginid")
    private String loginid;

    @Column(name = "u_passwd")
    private String passwd;

    @Column(name = "u_name")
    private String name;

    @Column(name = "u_level")
    private int level;

    @Column(name = "u_hp")
    private String hp;

    @Column(name = "u_email")
    private String email;

    @Column(name = "u_grade")
    private int grade;

    @Column(name = "u_status")
    private int status;

    @Column(name = "u_company")
    private Long company;

    @Column(name = "u_date")
    @CreationTimestamp
    private Timestamp date;
}
