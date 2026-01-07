drop view aptuserlist_vw;
create view aptuserlist_vw as select user_tb.*, au_user AS u_user, au_apt AS u_apt, au_level AS u_aptlevel, concat(ag_name, ' ', a_name) as u_aptname, a_status as u_aptstatus from user_tb, aptuser_tb, apt_tb, aptgroup_tb where u_id = au_user and au_apt = a_id and a_aptgroup = ag_id;
