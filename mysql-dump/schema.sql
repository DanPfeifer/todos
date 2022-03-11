CREATE DATABASE `todos`;

CREATE TABLE `todo` 
	id INT NOT NULL AUTO_INCREMENT,
  	title VARCHAR(256) NOT NULL,
  	message VARCHAR(256) NOT NULL,
   	insert_date DATETIME,
   	modified_date DATETIME,
   	due_date DATETIME,
   	user_id INT,
   	status TINYINT,
   	priority TINYINT
   PRIMARY KEY ( id );

 CREATE TABLE `user` 
	id INT NOT NULL AUTO_INCREMENT,
  	email VARCHAR(64) NOT NULL,
  	password VARCHAR(64) NOT NULL,
  	first_name VARCHAR(64) NOT NULL,
  	last_name VARCHAR(64) NOT NULL,
  	token VARCHAR(256) NOT NULL,
   	insert_date DATETIME,
   	modified_date DATETIME,
   	status TINYINT
   PRIMARY KEY ( id );

 INSERT INTO `user` (email, password, first_name, last_name, insert_date, modified_date, status)
 VALUES ('test@test.com', '123', 'Test', 'Testerson', NOW(), NOW(), 1);