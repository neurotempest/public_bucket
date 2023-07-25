

create table simple (
  id bigint not null auto_increment,

  a int,
  b char(255),
  c bool,

  primary key (id)
);

create table conversion (
  id bigint not null auto_increment,

  vc varchar(255),
  b blob,

  primary key(id)
);
