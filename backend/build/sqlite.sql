CREATE TABLE apt_tb (
  a_id integer,
  a_name text,
  a_status integer,
  a_date text
);

CREATE TABLE aptuser_tb (
  au_id integer,
  au_apt integer,
  au_user integer,
  au_level integer,
  au_date text
);

CREATE TABLE data_tb (
  d_id integer primary key AUTOINCREMENT,
  d_apt integer,
  d_image integer,
  d_imagetype integer,
  d_user integer,
  d_type integer,
  d_x real,
  d_y real,
  d_point text,
  d_number integer,
  d_group integer,
  d_name text,
  d_content text,
  d_width real,
  d_length real,
  d_count integer,
  d_progress text,
  d_remark text,
  d_imagename text,
  d_filename text,
  d_memo text,
  d_date text
);

CREATE TABLE findpasswd_tb (
  fp_id integer primary key AUTOINCREMENT,
  fp_user integer,
  fp_link text,
  fp_date text
);

CREATE TABLE image_tb (
  i_id integer,
  i_apt integer,
  i_name text,
  i_level integer,
  i_parent integer,
  i_last integer,
  i_title text,
  i_type integer,
  i_filename text,
  i_file1 text,
  i_file2 text,
  i_file3 text,
  i_file4 text,
  i_file5 text,
  i_date text
);

CREATE TABLE sendmail_tb (
  sm_id integer primary key AUTOINCREMENT,
  sm_level integer,
  sm_title text,
  sm_content text,
  sm_status integer,
  sm_date text
);

CREATE TABLE sendsms_tb (
  ss_id integer primary key AUTOINCREMENT,
  ss_level integer,
  ss_content text,
  ss_status integer,
  ss_date text
);


CREATE TABLE status_tb (
  s_id integer primary key AUTOINCREMENT,
  s_name text,
  s_type integer,
  s_content text,
  s_etc text,
  s_date text
);


CREATE TABLE sync_tb (
  s_id integer primary key AUTOINCREMENT,
  s_image integer,
  s_date text
);


CREATE TABLE user_tb (
  u_id integer,
  u_loginid text,
  u_passwd text,
  u_name text,
  u_level integer,
  u_hp text,
  u_email text,
  u_date text
);



create view aptuserlist_vw as select user_tb.*, au_user AS u_user, au_apt AS u_apt, au_level AS u_aptlevel, a_name as u_aptname from user_tb, aptuser_tb, apt_tb where u_id = au_user and au_apt = a_id;
