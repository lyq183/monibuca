CREATE DATABASE Monibuca;

CREATE TABLE department(
	d_id INT PRIMARY KEY AUTO_INCREMENT,
	d_name VARCHAR(30) NOT NULL,
	d_manager_id INT NOT NULL,
	d_description VARCHAR(100)
);

CREATE TABLE users(
	id INT PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(30) NOT NULL UNIQUE,
	PASSWORD VARCHAR(30) NOT NULL,
	department_id INT NOT NULL,
	FOREIGN KEY(department_id) REFERENCES department(d_id)
);

CREATE TABLE project(
	p_id INT PRIMARY KEY AUTO_INCREMENT,
	p_name VARCHAR(30),
	p_uid INT NOT NULL,
	p_department INT NOT NULL,
	FOREIGN KEY(p_uid) REFERENCES users(id),
	FOREIGN KEY(p_department) REFERENCES department(d_id)
);

CREATE TABLE sessions(
	session_id VARCHAR(100) PRIMARY KEY,
	permissions INT NOT NULL,
	user_id INT NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id)
);

create table admin(
	id INT PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(30) NOT NULL,
	PASSWORD VARCHAR(30) NOT NULL,
	session_id VARCHAR(100)
);
INSERT INTO admin(adminname,PASSWORD) VALUES ("admin","admin")
INSERT INTO department(d_name,d_manager_id,d_description) VALUES ("教学楼",3,"教学区")
INSERT INTO department(d_name,d_manager_id,d_description) VALUES ("实验室",4,"做实验")
INSERT INTO users(username,PASSWORD,department_id) VALUES ("12345","12345",1)
INSERT INTO project(p_name,p_uid,p_department,P_configtName) VALUES ("教室101",1,1,"c101.toml")
INSERT INTO project(p_name,p_uid,p_department,P_configtName) VALUES ("教室202",1,1,"c202.toml")