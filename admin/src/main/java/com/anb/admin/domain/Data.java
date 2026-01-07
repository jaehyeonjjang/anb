package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="data_tb")
public class Data {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "d_id")
    private Long id;

    @Column(name = "d_apt")
    private Long apt;

    @Column(name = "d_image")
    private Long image;

    @Column(name = "d_imagetype")
    private int imagetype;

    @Column(name = "d_user")
    private Long user;

    @Column(name = "d_type")
    private int type;    

    @Column(name = "d_x")
    private double x;

    @Column(name = "d_y")
    private double y;

    @Column(name = "d_point")
    private String point;

    @Column(name = "d_number")
    private int number;

    @Column(name = "d_group")
    private int group;

    @Column(name = "d_name")
    private String name;

    @Column(name = "d_fault")
    private String fault;

    @Column(name = "d_content")
    private String content;

    @Column(name = "d_width")
    private double width;

    @Column(name = "d_length")
    private double length;

    @Column(name = "d_count")
    private String count;

    @Column(name = "d_progress")
    private String progress;
    
    @Column(name = "d_remark")
    private String remark;

    @Column(name = "d_imagename")
    private String imagename;

    @Column(name = "d_filename")
    private String filename;

    @Column(name = "d_memo")
    private String memo;

    @Column(name = "d_report")
    private int report;

    @Column(name = "d_usermemo")
    private String usermemo;

    @Column(name = "d_aptmemo")
    private String aptmemo;

    @Column(name = "d_date")
    @CreationTimestamp
    private Timestamp date;
}
