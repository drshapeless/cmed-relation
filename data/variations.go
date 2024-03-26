package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Variation struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateVariation(variation *Variation, db DB) error {
	q := `INSERT INTO variations (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, variation.Content).Scan(&variation.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetVariation(id int64, db DB) (*Variation, error) {
	q := `SELECT * FROM variations WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	variation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Variation])
	if err != nil {
		return nil, err
	}

	return &variation, nil
}

func GetVariationByContent(content string, db DB) (*Variation, error) {
	q := `SELECT * FROM variations WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	variation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Variation])
	if err != nil {
		return nil, err
	}

	return &variation, nil
}

func UpdateVariation(variation *Variation, db DB) error {
	q := `UPDATE variations
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, variation.Content, variation.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteVariation(id int64, db DB) error {
	q := `DELETE FROM variations WHERE id = $1`

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
