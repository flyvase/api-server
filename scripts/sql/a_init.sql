create database mr_president;

create user harvest identified by 'FM5fX|jt(gdv-N%L';

grant select,
  insert,
  update on mr_president.* to harvest;