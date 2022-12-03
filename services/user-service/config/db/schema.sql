CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(50) NOT NULL,
    surname varchar(50) NOT NULL,
    email varchar(72) NOT NULL UNIQUE
);