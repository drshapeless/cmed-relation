package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Alternative struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateAlternative(alternative *Alternative, db DB) error {
	q := `INSERT INTO alternatives (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, alternative.Content).Scan(&alternative.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetAlternative(id int64, db DB) (*Alternative, error) {
	q := `SELECT * FROM alternatives WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	alternative, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Alternative])
	if err != nil {
		return nil, err
	}

	return &alternative, nil
}

func GetAlternativeByContent(content string, db DB) (*Alternative, error) {
	q := `SELECT * FROM alternatives WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	alternative, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Alternative])
	if err != nil {
		return nil, err
	}

	return &alternative, nil
}

func UpdateAlternative(alternative *Alternative, db DB) error {
	q := `UPDATE alternatives
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, alternative.Content, alternative.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAlternative(id int64, db DB) error {
	q := `DELETE FROM alternatives WHERE id = $1`

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
