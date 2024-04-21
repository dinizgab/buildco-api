SELECT c.id, c.name, c.email, c.phone, c.created_at, r.id, r.grade, r.comment, r.created_at
FROM company c
LEFT JOIN rating r ON r.company_id = c.id
WHERE c.id = $1;
