package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Method struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateMethod(method *Method, db DB) error {
	q := `INSERT INTO methods (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, method.Content).Scan(&method.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetMethod(id int64, db DB) (*Method, error) {
	q := `SELECT * FROM methods WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	method, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Method])
	if err != nil {
		return nil, err
	}

	return &method, nil
}

func GetMethodByContent(content string, db DB) (*Method, error) {
	q := `SELECT * FROM methods WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	method, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Method])
	if err != nil {
		return nil, err
	}

	return &method, nil
}

func UpdateMethod(method *Method, db DB) error {
	q := `UPDATE methods
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, method.Content, method.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMethod(id int64, db DB) error {
	q := `DELETE FROM methods WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	result, err := db.Exec(ctx, q, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
