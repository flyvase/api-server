use mr_president;

create table users (
  id int unsigned primary key auto_increment,
  firebase_uid varchar(255) not null unique,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  deleted_at datetime
);