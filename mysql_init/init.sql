create database umr;
create table if not exists umr.admin(
    id varchar(15) primary key,
    password varchar(255)
);