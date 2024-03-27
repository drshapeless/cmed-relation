package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Mechanism struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateMechanism(mechanism *Mechanism, db DB) error {
	q := `INSERT INTO mechanisms (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, mechanism.Content).Scan(&mechanism.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetMechanism(id int64, db DB) (*Mechanism, error) {
	q := `SELECT * FROM mechanisms WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	mechanism, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Mechanism])
	if err != nil {
		return nil, err
	}

	return &mechanism, nil
}

func GetMechanismByContent(content string, db DB) (*Mechanism, error) {
	q := `SELECT * FROM mechanisms WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	mechanism, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Mechanism])
	if err != nil {
		return nil, err
	}

	return &mechanism, nil
}

func UpdateMechanism(mechanism *Mechanism, db DB) error {
	q := `UPDATE mechanisms
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, mechanism.Content, mechanism.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteMechanism(id int64, db DB) error {
	q := `DELETE FROM mechanisms WHERE id = $1`

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
