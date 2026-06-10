CREATE DATABASE test_db;

CREATE TABLE user (id INT, name TEXT);

INSERT INTO user (id, name) VALUES (one, John);
INSERT INTO user (id, name) VALUES (two, Doe);

SELECT * FROM user;