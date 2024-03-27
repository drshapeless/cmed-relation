package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Usage struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateUsage(usage *Usage, db DB) error {
	q := `INSERT INTO usages (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, usage.Content).Scan(&usage.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetUsage(id int64, db DB) (*Usage, error) {
	q := `SELECT * FROM usages WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	usage, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Usage])
	if err != nil {
		return nil, err
	}

	return &usage, nil
}

func GetUsageByContent(content string, db DB) (*Usage, error) {
	q := `SELECT * FROM usages WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	usage, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Usage])
	if err != nil {
		return nil, err
	}

	return &usage, nil
}

func UpdateUsage(usage *Usage, db DB) error {
	q := `UPDATE usages
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, usage.Content, usage.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsage(id int64, db DB) error {
	q := `DELETE FROM usages WHERE id = $1`

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
