CREATE TABLE users
(
  id         VARCHAR(45) PRIMARY KEY ,
  email      VARCHAR(255) NOT NULL UNIQUE,
  password   BYTEA NOT NULL,
  ip_address VARCHAR(45),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
