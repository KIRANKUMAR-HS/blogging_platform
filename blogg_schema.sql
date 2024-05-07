
-- Create a new PostgreSQL database
CREATE DATABASE blogging;

-- Table to store users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);

-- Table to store blog posts
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    author_id INT REFERENCES users(id), -- Foreign key to users
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
