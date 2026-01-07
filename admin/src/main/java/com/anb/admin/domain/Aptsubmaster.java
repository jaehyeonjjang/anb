package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="aptsubmaster_tb")
public class Aptsubmaster {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "as_id")
    private Long id;

    @Column(name = "as_apt")
    private Long apt;

    @Column(name = "as_user")
    private Long user;

    @Column(name = "as_level")
    private int level;

    @Column(name = "as_company")
    private Long company;

    @Column(name = "as_date")
    @CreationTimestamp
    private Timestamp date;
}
