package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="company_tb")
public class Company {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "c_id")
    private Long id;

    @Column(name = "c_name")
    private String name;

    @Column(name = "c_ceo")
    private String ceo;

    @Column(name = "c_logo")
    private String logo;

    @Column(name = "c_stamp")
    private String stamp;

    @Column(name = "c_contractstartdate")
    private Timestamp contractstartdate;

    @Column(name = "c_contractenddate")
    private Timestamp contractenddate;

    @Column(name = "c_status")
    private int status;

    @Column(name = "c_date")
    @CreationTimestamp
    private Timestamp date;
}
