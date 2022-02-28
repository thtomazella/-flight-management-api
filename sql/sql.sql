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


DROP TABLE IF EXISTS aeroporto;

CREATE TABLE  aeroporto(
    id       	INT           auto_increment primary key,
	nome      	VARCHAR(200)  NOT NULL, 
	sigla   	VARCHAR(20)   NOT NULL,
    inclusion 	TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;

DROP TABLE IF EXISTS aeronave;

CREATE TABLE  aeronave(
    id       	INT           auto_increment primary key,
	nome      	VARCHAR(200)  NOT NULL, 
	prefixo   	VARCHAR(20)   NOT NULL,
	custo   	DOUBLE(10,2)  DEFAULT '0.00',
	preco   	DOUBLE(10,2)  DEFAULT '0.00',
    inclusion 	TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;

DROP TABLE IF EXISTS tipoequipamento;

CREATE TABLE  tipoequipamento(
    id       	INT           auto_increment primary key,
	nome      	VARCHAR(200)  NOT NULL, 
    inclusion 	TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;

DROP TABLE IF EXISTS tipo_voo;

CREATE TABLE  tipo_voo(
    id       	INT           auto_increment primary key,
	nome      	VARCHAR(200)  NOT NULL, 
    inclusion 	TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;

DROP TABLE IF EXISTS tipo_instrucao;

CREATE TABLE  tipo_instrucao(
    id       	INT           auto_increment primary key,
	nome      	VARCHAR(200)  NOT NULL, 
    inclusion 	TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;