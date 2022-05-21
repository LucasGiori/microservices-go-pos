CREATE TABLE "ticket"
(
    id UUID PRIMARY KEY NOT NULL,
    order_id UUID NOT NULL,
    description TEXT NOT NULL,
    date_time timestamp DEFAULT NOW(),
    email VARCHAR(254) NOT NULL,
    status VARCHAR(100) NOT NULL
);