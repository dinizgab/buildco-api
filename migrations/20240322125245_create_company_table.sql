-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS company (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO company (id, name, phone, email)
VALUES ('8db46e78-bf5b-46fb-8768-7e1fc457e5a7', 'Test co.1', 'testco1@gmail.com', '1234-1234');

INSERT INTO company (id, name, phone, email)
VALUES ('124f7323-ee68-4eb6-9509-84eb966cc5cf', 'Test co.2', 'testco2@gmail.com', '4321-4321');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS company;
-- +goose StatementEnd
