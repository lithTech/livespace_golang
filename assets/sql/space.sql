-- +migrate Up
create table planet(
    id serial primary key, 
    title text,
    population bigint,
    planet_type smallint,
    version int default 1
);

insert into planet (title, population, planet_type) values ('Earth', 23232009, 4);
insert into planet (title, population, planet_type) values ('Mercury', 0, 1);
insert into planet (title, population, planet_type) values ('Pluto', 0, 1);
insert into planet (title, population, planet_type) values ('Mars', 23232009, 2);