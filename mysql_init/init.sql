create database umr;
create table if not exists umr.admin(
    id varchar(15) primary key,
    password varchar(255)
);
create table if not exists umr.app(
    id varchar(15) primary key,
    name varchar(255) not null,
    description varchar(255) not null,
    consumerkey char(32) not null unique,
    consumersecret char(32) not null,
    callback varchar(255) not null,
    owner varchar(15)
);