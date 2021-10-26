CREATE TABLE IF NOT EXISTS cards(
    id varchar(50) PRIMARY KEY,
    label varchar(50),
    type int,
    content varchar(5000),
    created int,
    modified int
);