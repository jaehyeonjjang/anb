package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="companyuser_tb")
public class Companyuser {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "cu_id")
    private Long id;

    @Column(name = "cu_company")
    private Long company;

    @Column(name = "cu_user")
    private Long user;

    @Column(name = "cu_date")
    @CreationTimestamp
    private Timestamp date;
}
