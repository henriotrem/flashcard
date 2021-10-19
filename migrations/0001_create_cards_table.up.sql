CREATE TABLE IF NOT EXISTS cards(
    id serial NOT NULL PRIMARY KEY,
    label varchar(50),
    type int,
    content varchar(5000),
    created int,
    modified int
);