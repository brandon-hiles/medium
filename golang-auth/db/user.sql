CREATE TABLE users (
    user_id INT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted BOOLEAN DEFAULT FALSE,
    first_name VARCHAR ( 50 ),
    last_name VARCHAR ( 50 ),
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    password VARCHAR ( 255 ) NOT NULL,
    PRIMARY KEY (user_id)
);