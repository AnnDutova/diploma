Create user pguser with password 'pgpwd';

create database testdb;

GRANT CONNECT on database "testdb" to pguser;

create table if not exists "project" (
     id serial primary key,
     title text not null unique
);

create table if not exists "author" (
    id serial primary key,
    name text not null unique
);

GRANT SELECT, INSERT, UPDATE, DELETE on table "project","author" to pguser;

