CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar (50) UNIQUE NOT NULL, 
    email varchar(255) UNIQUE NOT NULL,
    firstName varchar(255) NOT NULL,
    lastName varchar(255) NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_Active BOOLEAN NOT NULL DEFAULT TRUE
);