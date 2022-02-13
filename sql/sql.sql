-- CREATE A NEW USER
CREATE USER 'nomeUser'@'localhost' IDENTIFIED BY '***KEY***';

-- PRIVILEGE
 GRANT ALL PRIVILEGES ON * . * TO 'nomeUser'@'localhost';

-- CREATE A BASE
CREATE DATABASE IF NOT EXISTS flight_control;

USE flight_control;

-- CREATE TABLES

DROP TABLE IF EXISTS usuario;

CREATE TABLE  usuario(
    id       	INT           auto_increment primary key,
	nome      	VARCHAR(200)  NOT NULL, 
	type_user   VARCHAR(200)  NOT NULL, 
    nick      	VARCHAR(20)   NOT NULL, 
	cpf       	VARCHAR(20)   DEFAULT NULL, 
	id_anac 	VARCHAR(10)   NOT NULL, 
	cma       	DATE NULL     DEFAULT NULL,
	address  	VARCHAR(200)  DEFAULT NULL,
	number   	VARCHAR(10)   DEFAULT NULL,
	district    VARCHAR(200)  DEFAULT NULL,
	id_city 	VARCHAR(5)    DEFAULT NULL,
	contact     VARCHAR(15)   DEFAULT NULL,
	cell   		VARCHAR(15)   DEFAULT NULL,
	email     	VARCHAR(200)  NOT NULL,
	senha    	VARCHAR(200)  NOT NULL, 
    inclusion 	TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;

