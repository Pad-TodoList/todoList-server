CREATE TABLE todolistUser (
    id SERIAL,
    nickname VARCHAR(100) NOT NULL UNIQUE,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);
CREATE TABLE userToken (
    userid VARCHAR(100) NOT NULL,
    token VARCHAR(100) NOT NULL
);