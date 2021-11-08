create database mr_president character set utf8mb4;

create user harvest identified by 'FM5fX|jt(gdv-N%L';

grant select,
  insert,
  update on mr_president.* to harvest;