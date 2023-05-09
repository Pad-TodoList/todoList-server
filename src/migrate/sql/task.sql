CREATE TABLE task (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    startDate VARCHAR(100),
    endDate VARCHAR(100),
    status VARCHAR(10) NOT NULL,
    userId VARCHAR(100) not null
);