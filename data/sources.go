package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Source struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
	Remark  string `db:"remark"`
}

func CreateSource(source *Source, db DB) error {
	q := `INSERT INTO sources (content, remark)
VALUES ($1, $2)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, source.Content, source.Remark).Scan(&source.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetSource(id int64, db DB) (*Source, error) {
	q := `SELECT * FROM sources WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	source, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Source])
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func UpdateSource(source *Source, db DB) error {
	q := `UPDATE sources
SET content = $1, remark = $2
WHERE id = $3`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, source.Content, source.Remark, source.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteSource(id int64, db DB) error {
	q := `DELETE FROM sources WHERE id = $1`

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
