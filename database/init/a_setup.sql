create database main_database character set utf8mb4;

create user api_server identified by 'FM5fX|jt(gdv-N%L';

grant select,
  insert,
  update on main_database.* to api_server;