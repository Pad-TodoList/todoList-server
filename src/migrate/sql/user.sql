USE todolist;
DROP TABLE IF EXISTS user;
CREATE TABLE todolistUser (
    id int primary key NOT NULL UNIQUE,
    nickname VARCHAR(100) NOT NULL UNIQUE,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100) NOT NULL
);
DROP TABLE IF EXISTS user;
CREATE TABLE userToken (
    userId VARCHAR(100) NOT NULL,
    token VARCHAR(100) NOT NULL
);