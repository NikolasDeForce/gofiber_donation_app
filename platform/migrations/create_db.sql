DROP DATABASE IF EXISTS donationapp_rest;
CREATE DATABASE donationapp_rest;

SET TIMEZONE="Europe/Moscow";

\c donationapp_rest

-- Create register user table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    login VARCHAR (255) NOT NULL,
    email VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL
);

-- Create donates table
CREATE TABLE donates (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    loginwhodonate VARCHAR (255) NOT NULL,
    logintodonate VARCHAR (255) NOT NULL,
    message VARCHAR (255) NOT NULL,
    summary INT NOT NULL
);
-- Create admin user
INSERT INTO users(created_at, login, email, password) VALUES ('2024-01-01 00:00:00', 'admin', 'admin@mail.ru', 'admin');

-- Add indexes
CREATE INDEX notadmin_users ON users (email) WHERE login != 'admin' AND password != 'admin';