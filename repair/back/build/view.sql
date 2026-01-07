drop view totalreport_vw;
create view totalreport_vw as
select b_apt as b_id, b_apt, b_topcategory, b_subcategory, b_category, b_standard, b_method, b_rate, b_lastdate, b_duedate, sum(b_count) as b_count, now() as b_date from breakdown_tb group by b_apt, b_topcategory, b_subcategory, b_category, b_standard, b_method, b_rate, b_lastdate, b_duedate;

drop view totalyearreport_vw;
create view totalyearreport_vw as
select b_apt as b_id, b_apt, b_topcategory, b_subcategory, b_category, b_standard, b_method, b_rate, b_duedate, sum(b_count) as b_count, now() as b_date from breakdown_tb group by b_apt, b_topcategory, b_subcategory, b_category, b_standard, b_method, b_rate, b_duedate;


drop view standardlist_vw;
create view standardlist_vw as
select standard_tb.*, a.c_parent as s_subcategory, a.c_order as s_categoryorder, b.c_parent as s_topcategory from standard_tb, category_tb a, category_tb b where s_category = a.c_id and a.c_parent = b.c_id;

drop view yearreport_vw;
create view yearreport_vw as
select b_apt as b_id, b_apt, b_topcategory, b_subcategory, b_category, b_standard, b_method, b_rate, b_duedate, sum(b_count) as b_count, now() as b_date from breakdown_tb group by b_apt, b_topcategory, b_subcategory, b_category, b_standard, b_method, b_rate, b_duedate;


drop view repairlist_vw;
create view repairlist_vw as
select apt_tb.*, r_id as a_repairid, r_type as a_repairtype, r_reportdate as a_reportdate, r_date as a_repairdate, r_info1 as a_info1, if (r_type <> 3 and r_status = 1, 1, 2) as a_status from apt_tb, repair_tb where a_id = r_apt;


drop view periodicdatabackup_vw;
create view periodicdatabackup_vw as
select pd_periodic as pd_id, pd_date, pd_blueprint, count(*) as pd_count from periodicdata_tb where pd_periodic < 0 group by pd_periodic, pd_date, pd_blueprint order by pd_periodic;


drop view aptrepairlist_vw;
create view aptrepairlist_vw as 
select a_id, a_name, a_tel, a_fax, a_testdate, a_email, a_personalemail, a_zip, a_address, a_address2, a_completeyear, a_type, a_flatcount, a_familycount, a_floor, a_fmsloginid, a_fmspasswd, GROUP_CONCAT(r_reportdate) as a_reportdate from apt_tb, repair_tb where a_id = r_apt and r_type in (1, 2) and r_reportdate <> '' group by a_id, a_name, a_tel, a_fax, a_testdate, a_email, a_personalemail, a_zip, a_address, a_address2, a_completeyear, a_type, a_flatcount, a_familycount, a_floor, a_fmsloginid, a_fmspasswd;

drop view aptlist_vw;
create view aptlist_vw as 
select a_id, a_name, a_tel, a_fax, a_testdate, a_email, a_personalemail, a_zip, a_address, a_address2, a_completeyear, a_type, a_flatcount, a_familycount, a_floor, a_fmsloginid, a_fmspasswd, ifnull(GROUP_CONCAT(r_reportdate), '') as a_repairdate, ifnull(GROUP_CONCAT(d_reportdate), '') as a_periodicdate from apt_tb left outer join repair_tb on a_id = r_apt and r_type in (1, 2) left outer join periodic_tb on a_id = d_apt group by a_id, a_name, a_tel, a_fax, a_testdate, a_email, a_personalemail, a_zip, a_address, a_address2, a_completeyear, a_type, a_flatcount, a_familycount, a_floor, a_fmsloginid, a_fmspasswd;
