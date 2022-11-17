CREATE TABLE users
(
    id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR(250),
    last_name VARCHAR(250),
    username VARCHAR(70),
    phone VARCHAR(20),
    email VARCHAR(100),
    password VARCHAR(15),
    address VARCHAR(150),
    gender VARCHAR(6),
    role VARCHAR(50),
    postalcode INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)