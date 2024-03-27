package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Cause struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateCause(cause *Cause, db DB) error {
	q := `INSERT INTO causes (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, cause.Content).Scan(&cause.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetCause(id int64, db DB) (*Cause, error) {
	q := `SELECT * FROM causes WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	cause, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Cause])
	if err != nil {
		return nil, err
	}

	return &cause, nil
}

func GetCauseByContent(content string, db DB) (*Cause, error) {
	q := `SELECT * FROM causes WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	cause, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Cause])
	if err != nil {
		return nil, err
	}

	return &cause, nil
}

func UpdateCause(cause *Cause, db DB) error {
	q := `UPDATE causes
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, cause.Content, cause.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteCause(id int64, db DB) error {
	q := `DELETE FROM causes WHERE id = $1`

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
