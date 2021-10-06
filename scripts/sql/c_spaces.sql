-- Only for local development env
use mr_president;

create table spaces (
  id int unsigned primary key auto_increment,
  name varchar(150) not null
);

-- get all statement example
-- select * from spaces;
