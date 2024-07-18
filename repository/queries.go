package repository

const (
	InsertQuery = `INSERT INTO "users" (username, email, password_hash, firstname, lastname, birthday, country) VALUES ($1, $2, $3, $4, $5, $6, $7)`
)
