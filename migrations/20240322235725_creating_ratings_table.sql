-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rating (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    grade INTEGER NOT NULL,
    comment TEXT NOT NULL,
    company_id UUID NOT NULL REFERENCES company(id),
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rating;
-- +goose StatementEnd
