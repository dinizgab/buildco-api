-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS rating (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    grade INTEGER NOT NULL,
    comment TEXT NOT NULL,
    company_id UUID NOT NULL REFERENCES company(id),
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO rating (grade, comment, company_id)
VALUES (1, 'Worst company ever!', '124f7323-ee68-4eb6-9509-84eb966cc5cf'),
(5, 'Best company ever!', '124f7323-ee68-4eb6-9509-84eb966cc5cf'),
(3, 'Mid company ever!', '124f7323-ee68-4eb6-9509-84eb966cc5cf');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS rating;
-- +goose StatementEnd
