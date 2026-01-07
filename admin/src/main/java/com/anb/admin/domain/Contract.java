package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="contract_tb")
public class Contract {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "co_id")
    private Long id;

    @Column(name = "co_company")
    private Long company;

    @Column(name = "co_status")
    private int status;

    @Column(name = "co_contractstartdate")
    private Timestamp contractstartdate;

    @Column(name = "co_contractenddate")
    private Timestamp contractenddate;

    @Column(name = "co_date")
    @CreationTimestamp
    private Timestamp date;
}
