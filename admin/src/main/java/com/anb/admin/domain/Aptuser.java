package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="aptuser_tb")
public class Aptuser {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "au_id")
    private Long id;

    @Column(name = "au_apt")
    private Long apt;

    @Column(name = "au_user")
    private Long user;

    @Column(name = "au_level")
    private int level;

    @Column(name = "au_company")
    private Long company;

    @Column(name = "au_date")
    @CreationTimestamp
    private Timestamp date;
}
