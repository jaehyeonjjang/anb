package com.anb.admin.domain;

import lombok.Getter;
import lombok.Setter;
import org.hibernate.annotations.CreationTimestamp;

import javax.persistence.*;
import java.sql.Timestamp;

@Getter
@Setter
@Entity
@Table(name="report_tb")
public class Report {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "r_id")
    private Long id;

    @Column(name = "r_apt")
    private Long apt;

    @Column(name = "r_image")
    private Long image;

    @Column(name = "r_status")
    private Long status;

    @Column(name = "r_date")
    @CreationTimestamp
    private Timestamp date;
}
