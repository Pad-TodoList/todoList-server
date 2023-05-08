USE todolist;
DROP TABLE IF EXISTS task;
CREATE TABLE task (
    id int primary key NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    startDate VARCHAR(100),
    endDate VARCHAR(100),
    status VARCHAR(10) NOT NULL,
    userId VARCHAR(100) not null
);