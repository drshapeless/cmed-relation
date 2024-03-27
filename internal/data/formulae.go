package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Formulae struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateFormulae(formulae *Formulae, db DB) error {
	q := `INSERT INTO formulae (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, formulae.Content).Scan(&formulae.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetFormulae(id int64, db DB) (*Formulae, error) {
	q := `SELECT * FROM formulae WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	formulae, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Formulae])
	if err != nil {
		return nil, err
	}

	return &formulae, nil
}

func GetFormulaeByContent(content string, db DB) (*Formulae, error) {
	q := `SELECT * FROM formulae WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	formulae, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Formulae])
	if err != nil {
		return nil, err
	}

	return &formulae, nil
}

func UpdateFormulae(formulae *Formulae, db DB) error {
	q := `UPDATE formulae
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, formulae.Content, formulae.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteFormulae(id int64, db DB) error {
	q := `DELETE FROM formulae WHERE id = $1`

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
