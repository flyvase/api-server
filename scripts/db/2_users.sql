-- Only for local development env
create table users (
  id int unsigned primary key auto_increment,
  firebase_uid varchar(255) not null unique,
  first_name varchar(100) not null,
  last_name varchar(100) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  deleted_at datetime
);

-- insert statement example
-- insert into users (firebase_uid, first_name, last_name)
-- values (
--     "cUOv9IiMsAgpzTzoUjkTpwW9Ilr2",
--     "Shunei",
--     "Hayakawa"
--   );