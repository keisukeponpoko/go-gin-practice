USE test;

DROP TABLE IF EXISTS user;

CREATE TABLE user (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user varchar(20) NOT NULL
);

INSERT INTO user (id, user) VALUES (1, "Taro");
