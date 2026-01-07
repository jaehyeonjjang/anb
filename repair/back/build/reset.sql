
truncate area_tb;
truncate breakdown_tb;
delete from category_tb where c_apt <> -1;
delete from dong_tb where d_apt <> -1;
truncate history_tb;
truncate repair_tb;
delete from standard_tb where s_apt <> -1;

