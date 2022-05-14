CREATE TABLE "user"
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    login VARCHAR(200) NOT NULL,
    password VARCHAR(72) NOT NULL,
    CONSTRAINT user_uk UNIQUE(login)
);