CREATE DATABASE test_db;

CREATE TABLE user (id INT, name TEXT);

INSERT INTO user (id, name) VALUES (1, John);
INSERT INTO user (id, name) VALUES (2, Doe);

SELECT * FROM user;