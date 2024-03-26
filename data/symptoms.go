package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Symptom struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateSymptom(symptom *Symptom, db DB) error {
	q := `INSERT INTO symptoms (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, symptom.Content).Scan(&symptom.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetSymptom(id int64, db DB) (*Symptom, error) {
	q := `SELECT * FROM symptoms WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	symptom, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Symptom])
	if err != nil {
		return nil, err
	}

	return &symptom, nil
}

func GetSymptomByContent(content string, db DB) (*Symptom, error) {
	q := `SELECT * FROM symptoms WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	symptom, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Symptom])
	if err != nil {
		return nil, err
	}

	return &symptom, nil
}

func UpdateSymptom(symptom *Symptom, db DB) error {
	q := `UPDATE symptoms
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, symptom.Content, symptom.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteSymptom(id int64, db DB) error {
	q := `DELETE FROM symptoms WHERE id = $1`

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
