SELECT c.id, c.name, c.email, c.phone, c.created_at, r.grade, r.comment
FROM company c
LEFT JOIN rating r ON r.company_id = c.id
WHERE c.id = $1;
