-- CREATE A NEW USER
CREATE USER 'nomeUser'@'localhost' IDENTIFIED BY '***KEY***';

-- PRIVILEGE
 GRANT ALL PRIVILEGES ON * . * TO 'nomeUser'@'localhost';

-- CREATE A BASE
CREATE DATABASE IF NOT EXISTS flight_control;

USE flight_control;

-- CREATE TABLES

DROP TABLE IF EXISTS users;

CREATE TABLE  users(
    id       INT         auto_increment primary key,
	nome      VARCHAR(200)  NOT NULL, 
    nick      VARCHAR(20)   NOT NULL, 
    senha     VARCHAR(200)  NOT NULL, 
	cpf       VARCHAR(20)   DEFAULT NULL, 
	cAnac     VARCHAR(10)   NOT NULL, 
	cma       DATE NULL     DEFAULT NULL,
	endereco  VARCHAR(200)  DEFAULT NULL,
	numero    VARCHAR(10)   DEFAULT NULL,
	bairro    VARCHAR(200)  DEFAULT NULL,
	codCidade VARCHAR(5)    DEFAULT NULL,
	fone1     VARCHAR(15)   DEFAULT NULL,
	fone2     VARCHAR(15)   DEFAULT NULL,
	celular   VARCHAR(15)   DEFAULT NULL,
	email     VARCHAR(200)  DEFAULT NULL,
    criadoem TIMESTAMP   DEFAULT current_timestamp()
 )ENGINE=INNODB;

