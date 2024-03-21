DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'hexagonal') THEN
        CREATE DATABASE hexagonal;
    END IF;
END
$$;

-- Create Table Users
DROP TABLE IF EXISTS todo;
CREATE TABLE todo(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" VARCHAR(50) NOT NULL,
    "description" VARCHAR(50) NOT NULL,
    "done" BOOLEAN NOT NULL DEFAULT false
);