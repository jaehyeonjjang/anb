package com.anb.admin.domain;

import java.sql.Timestamp;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import org.hibernate.annotations.CreationTimestamp;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@Entity
@Table(name = "statuscategory_tb")
public class Statuscategory {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "sc_id")
    private Long id;

    @Column(name = "sc_name")
    private String name;

    @Column(name = "sc_type")
    private int type;

    @Column(name = "sc_floortype")
    private int floortype;

    @Column(name = "sc_order")
    private int order;

    @Column(name = "sc_company")
    private Long company;

    @Column(name = "sc_date")
    @CreationTimestamp
    private Timestamp date;
}
