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
VALUES ('244c423a-930f-42a3-837f-c99102d27339', 'Test co.1', 'testco1@gmail.com', '1234-1234')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS company;
-- +goose StatementEnd
