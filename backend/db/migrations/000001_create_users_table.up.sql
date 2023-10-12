CREATE TABLE IF NOT EXISTS users(
    id serial primary key,
    firstName VARCHAR (55) NOT NULL,
    lastName VARCHAR (55) NOT NULL,
    preferredName VARCHAR (20),
    email VARCHAR (55) UNIQUE,
    state VARCHAR (20),
    imageUrl VARCHAR (255),
    role VARCHAR (20),
    createdAt timestamp
);