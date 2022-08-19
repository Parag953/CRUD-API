DROP DATABASE IF EXISTS Users;
CREATE DATABASE IF NOT EXISTS Users;
use Users;

CREATE TABLE IF NOT EXISTS userinfo(
    id int not null AUTO_INCREMENT,
    name varchar(50) not null,
    contact_no char(20) not null,
    DOB DATE not null,
    Primary Key (id)
);

insert into userinfo
    (name, contact_no, DOB)
VALUES
    ('Parag' , '7073670256' ,'1998-06-24'),
    ('Pulkit', '8639587543' ,'1998-06-24'),
    ('Shorya', '1234567890' ,'1996-03-21');-
    