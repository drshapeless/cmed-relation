package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Herb struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateHerb(herb *Herb, db DB) error {
	q := `INSERT INTO herbs (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, herb.Content).Scan(&herb.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetHerb(id int64, db DB) (*Herb, error) {
	q := `SELECT * FROM herbs WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	herb, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Herb])
	if err != nil {
		return nil, err
	}

	return &herb, nil
}

func GetHerbByContent(content string, db DB) (*Herb, error) {
	q := `SELECT * FROM herbs WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	herb, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Herb])
	if err != nil {
		return nil, err
	}

	return &herb, nil
}

func UpdateHerb(herb *Herb, db DB) error {
	q := `UPDATE herbs
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, herb.Content, herb.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteHerb(id int64, db DB) error {
	q := `DELETE FROM herbs WHERE id = $1`

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
