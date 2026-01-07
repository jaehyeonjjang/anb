package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="status_tb")
public class Status {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "s_id")
    private Long id;

    @Column(name = "s_name")
    private String name;

    @Column(name = "s_statuscategory")
    private Long statuscategory;

    @Column(name = "s_type")
    private int type;
    
    @Column(name = "s_content")
    private String content;

    @Column(name = "s_etc")
    private String etc;

    @Column(name = "s_order")
    private int order;

    @Column(name = "s_company")
    private Long company;

    @Column(name = "s_date")
    @CreationTimestamp
    private Timestamp date;
}
