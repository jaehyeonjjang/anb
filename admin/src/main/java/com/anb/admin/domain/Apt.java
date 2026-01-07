package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="apt_tb")
public class Apt {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "a_id")
    private Long id;

    @Column(name = "a_aptgroup")
    private Long aptgroup;

    @Column(name = "a_name")
    private String name;

    @Column(name = "a_startdate")
    private String startdate;

    @Column(name = "a_enddate")
    private String enddate;

    @Column(name = "a_type")
    private int type;

    @Column(name = "a_master")
    private Long master;

    @Column(name = "a_status")
    private int status;

    @Column(name = "a_company")
    private Long company;

    @Column(name = "a_report")
    private int report;

    @Column(name = "a_report1")
    private int report1;

    @Column(name = "a_report2")
    private int report2;

    @Column(name = "a_report3")
    private int report3;

    @Column(name = "a_report4")
    private int report4;

    @Column(name = "a_report5")
    private int report5;

    @Column(name = "a_report6")
    private int report6;

    @Column(name = "a_summarytype")
    private int summarytype;

    @Column(name = "a_search")
    private String search;

    @Column(name = "a_user")
    private Long user;

    @Column(name = "a_updateuser")
    private Long updateuser;

    @Column(name = "a_date")
    @CreationTimestamp
    private Timestamp date;
}
