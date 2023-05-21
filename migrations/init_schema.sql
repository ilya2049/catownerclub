create table cat_owners (
  id int primary key,
  name text not null
);

create table cats (
  id int primary key,
  owner_id int not null references cat_owners (id),
  name text not null
);