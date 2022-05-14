CREATE TABLE customer
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(200) NOT NULL,
    email VARCHAR(200) NOT NULL
);