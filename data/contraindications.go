package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Contraindication struct {
	ID      int64  `db:"id"`
	Content string `db:"content"`
}

func CreateContraindication(contraindication *Contraindication, db DB) error {
	q := `INSERT INTO contraindications (content)
VALUES ($1)
RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.QueryRow(ctx, q, contraindication.Content).Scan(&contraindication.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetContraindication(id int64, db DB) (*Contraindication, error) {
	q := `SELECT * FROM contraindications WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	contraindication, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contraindication])
	if err != nil {
		return nil, err
	}

	return &contraindication, nil
}

func GetContraindicationByContent(content string, db DB) (*Contraindication, error) {
	q := `SELECT * FROM contraindications WHERE content = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, content)
	if err != nil {
		return nil, err
	}

	contraindication, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Contraindication])
	if err != nil {
		return nil, err
	}

	return &contraindication, nil
}

func UpdateContraindication(contraindication *Contraindication, db DB) error {
	q := `UPDATE contraindications
SET content = $1
WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, contraindication.Content, contraindication.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteContraindication(id int64, db DB) error {
	q := `DELETE FROM contraindications WHERE id = $1`

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
