create table user (
  id bigint primary key auto_increment,
  state bigint,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  first_name varchar(255),
  last_name varchar(255)
);
