package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Symptomx struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateSymptomx(symptomx *Symptomx, db DB) error {
	q := `INSERT INTO symptomx (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, symptomx.Content).Scan(&symptomx.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetSymptomx(id int64, db DB) (*Symptomx, error) {
	q := `SELECT * FROM symptomx WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	symptomx, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Symptomx])
	if err != nil {
		return nil, err
	}

	return &symptomx, nil
}

func GetSymptomxByContent(content string, db DB) (*Symptomx, error) {
	q := `SELECT * FROM symptomx WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	symptomx, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Symptomx])
	if err != nil {
		return nil, err
	}

	return &symptomx, nil
}

func UpdateSymptomx(symptomx *Symptomx, db DB) error {
	q := `UPDATE symptomx
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, symptomx.Content, symptomx.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteSymptomx(id int64, db DB) error {
	q := `DELETE FROM symptomx WHERE id = $1`

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
