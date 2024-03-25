INSERT INTO users (name, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING id, name, username, email;
