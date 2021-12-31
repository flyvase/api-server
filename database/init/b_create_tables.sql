use main;

create table space_owners (
  id int unsigned primary key auto_increment,
  name varchar(120) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  deleted_at datetime
);

create table spaces (
  id int unsigned primary key auto_increment,
  owner_id int unsigned,
  foreign key (owner_id) references space_owners(id) on delete cascade on update cascade,
  headline varchar(240) not null,
  access varchar(80),
  weekly_visitors int unsigned,
  -- based on ISO5218
  main_customers_sex char(1) not null default '0',
  min_main_customers_age tinyint unsigned,
  max_main_customers_age tinyint unsigned,
  daily_price int unsigned not null,
  website_url text,
  coordinate point not null srid 4326,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  deleted_at datetime
);

create table space_images (
  id int unsigned primary key auto_increment,
  space_id int unsigned,
  foreign key (space_id) references spaces(id) on delete cascade on update cascade,
  image_url text not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  deleted_at datetime
);

create table space_displays (
  id int unsigned primary key auto_increment,
  space_id int unsigned,
  foreign key (space_id) references spaces(id) on delete cascade on update cascade,
  image_url text not null,
  description varchar(120) not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  deleted_at datetime
);