package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;
import javax.persistence.Transient;

import com.anb.admin.domain.Apt;

@Getter
@Setter
@Entity
@Table(name="aptgroup_tb")
public class Aptgroup {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "ag_id")
    private Long id;

    @Column(name = "ag_company")
    private Long company;

    @Column(name = "ag_name")
    private String name;

    @Column(name = "ag_facility")
    private int facility;

    @Column(name = "ag_type")
    private int type;

    @Column(name = "ag_status")
    private int status;

    @Column(name = "ag_imagecategory")
    private String imagecategory;

    @Column(name = "ag_user")
    private Long user;

    @Column(name = "ag_updateuser")
    private Long updateuser;

    @Column(name = "ag_date")
    @CreationTimestamp
    private Timestamp date;
}
