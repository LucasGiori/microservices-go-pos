CREATE TABLE product
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    name VARCHAR(200) NOT NULL,
    value numeric(19, 2) NOT NULL
);