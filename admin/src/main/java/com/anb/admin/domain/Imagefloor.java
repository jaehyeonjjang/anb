package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="imagefloor_tb")
public class Imagefloor {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "if_id")
    private Long id;

    @Column(name = "if_image")
    private Long image;

    @Column(name = "if_name")
    private String name;

    @Column(name = "if_imagename")
    private String imagename;

    @Column(name = "if_target")
    private Long target;

    @Column(name = "if_date")
    @CreationTimestamp
    private Timestamp date;
}
