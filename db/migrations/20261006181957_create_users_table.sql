-- +goose Up
CREATE TABLE users(
    id  int PRIMARY KEY AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    created_At TIMESTAMP Default CURRENT_TIMESTAMP,
    updated_At TIMESTAMP Default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;
