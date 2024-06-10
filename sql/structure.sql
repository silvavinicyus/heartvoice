CREATE DATABASE IF NOT EXISTS heartvoice_dev;

USE heartvoice_dev;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null,
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  password text not null,
  createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

DROP TABLE IF EXISTS guilds;

CREATE TABLE guilds (
  id int auto_increment primary key,
  name varchar(50) not null unique,
  imageKey varchar(255) unique,
  description TEXT not null,
  createdAt timestamp default current_timestamp(),
  updatedAt timestamp default current_timestamp()
) ENGINE=INNODB;
