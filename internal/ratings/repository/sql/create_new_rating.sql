INSERT INTO rating (grade, comment, company_id)
VALUES ($1, $2, $3)
RETURNING id, grade, comment, company_id
