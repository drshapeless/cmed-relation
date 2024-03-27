package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Disease struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateDisease(disease *Disease, db DB) error {
	q := `INSERT INTO diseases (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, disease.Content).Scan(&disease.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetDisease(id int64, db DB) (*Disease, error) {
	q := `SELECT * FROM diseases WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	disease, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Disease])
	if err != nil {
		return nil, err
	}

	return &disease, nil
}

func GetDiseaseByContent(content string, db DB) (*Disease, error) {
	q := `SELECT * FROM diseases WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	disease, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Disease])
	if err != nil {
		return nil, err
	}

	return &disease, nil
}

func UpdateDisease(disease *Disease, db DB) error {
	q := `UPDATE diseases
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, disease.Content, disease.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteDisease(id int64, db DB) error {
	q := `DELETE FROM diseases WHERE id = $1`

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
