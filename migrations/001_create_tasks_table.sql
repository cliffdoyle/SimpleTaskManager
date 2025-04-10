-- This SQL script creates a table named 'tasks' with the following columns:
CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(100) NOT NULL DEFAULT 'pemding',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);