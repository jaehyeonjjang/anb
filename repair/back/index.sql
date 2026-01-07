create index standard_tb_category on standard_tb (s_category);

create index breakdown_tb_method on breakdown_tb (b_method);
create index breakdown_tb_category on breakdown_tb (b_category);

create index history_tb_category on history_tb (h_category);

create index review_tb_method on review_tb (re_method);
create index review_tb_category on review_tb (re_category);

create index adjust_tb_category on adjust_tb (aj_category);

create index breakdownhistory_tb_method on breakdownhistory_tb (bh_method);
create index breakdownhistory_tb_category on breakdownhistory_tb (bh_category);

create index category_tb_apt_parent on category_tb (c_apt, c_parent);
create index category_tb_apt_order on category_tb (c_apt, c_order);
create index category_tb_apt_name on category_tb (c_apt, c_name);
create index category_tb_apt_level_parent_name on category_tb (c_apt, c_level, c_parent, c_name);

update repair_tb set r_content10 = replace(r_content10, '300만원', '500만원') where r_content10 like '%300만원%';


ALTER TABLE `repair_tb` ADD `r_periodtype` INT NOT NULL AFTER `r_content11`;

ALTER TABLE `saving_tb` ADD `sa_etc` BIGINT NOT NULL AFTER `sa_saving`;
update saving_tb set sa_etc = 0;
