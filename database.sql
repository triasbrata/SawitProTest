CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE estate (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    width INT NOT NULL,
    "length" INT NOT NULL
);

CREATE TABLE tree (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    estate_id UUID,
    x INT NOT NULL,
    y INT NOT NULL,
    height INT NOT NULL
);
