CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar (50) UNIQUE NOT NULL, 
    password varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);