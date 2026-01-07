package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="work_tb")
public class Work {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "w_id")
    private Long id;

    @Column(name = "w_image")
    private Long image;

    @Column(name = "w_type")
    private int type;

    @Column(name = "w_content")
    private String content;

    @Column(name = "w_pos")
    private int pos;

    @Column(name = "w_date")
    @CreationTimestamp
    private Timestamp date;
}
