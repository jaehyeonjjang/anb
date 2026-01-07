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
@Table(name = "image_tb")
public class Image {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "i_id")
    private Long id;

    @Column(name = "i_apt")
    private Long apt;

    @Column(name = "i_name")
    private String name;

    @Column(name = "i_level")
    private int level;

    @Column(name = "i_parent")
    private Long parent;

    @Column(name = "i_last")
    private int last;

    @Column(name = "i_title")
    private String title;

    @Column(name = "i_type")
    private int type;

    @Column(name = "i_floortype")
    private int floortype;

    @Column(name = "i_filename")
    private String filename;

    @Column(name = "i_order")
    private int order;

    @Column(name = "i_standard")
    private int standard;
    
    @Column(name = "i_date")
    @CreationTimestamp
    private Timestamp date;
}
