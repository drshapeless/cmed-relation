package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type SourceDiseaseRelation struct {
	SourceID  int64 `db:"source_id"`
	DiseaseID int64 `db:"disease_id"`
}

func CreateSourceDiseaseRelation(sourceDiseaseRelation *SourceDiseaseRelation, db DB) error {
	q := `INSERT INTO source_disease_relation (source_id, disease_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceDiseaseRelation.SourceID, sourceDiseaseRelation.DiseaseID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceDiseaseRelationBySourceID(id, db DB) (*SourceDiseaseRelation, error) {
	q := `SELECT * FROM source_disease_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceDiseaseRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceDiseaseRelation])
	if err != nil {
		return nil, err
	}

	return &sourceDiseaseRelation, nil
}

func GetSourceDiseaseRelationByDiseaseID(id, db DB) (*SourceDiseaseRelation, error) {
	q := `SELECT * FROM source_disease_relation WHERE disease_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceDiseaseRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceDiseaseRelation])
	if err != nil {
		return nil, err
	}

	return &sourceDiseaseRelation, nil
}

type SourceCauseRelation struct {
	SourceID int64 `db:"source_id"`
	CauseID  int64 `db:"cause_id"`
}

func CreateSourceCauseRelation(sourceCauseRelation *SourceCauseRelation, db DB) error {
	q := `INSERT INTO source_cause_relation (source_id, cause_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceCauseRelation.SourceID, sourceCauseRelation.CauseID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceCauseRelationBySourceID(id, db DB) (*SourceCauseRelation, error) {
	q := `SELECT * FROM source_cause_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceCauseRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceCauseRelation])
	if err != nil {
		return nil, err
	}

	return &sourceCauseRelation, nil
}

func GetSourceCauseRelationByCauseID(id, db DB) (*SourceCauseRelation, error) {
	q := `SELECT * FROM source_cause_relation WHERE cause_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceCauseRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceCauseRelation])
	if err != nil {
		return nil, err
	}

	return &sourceCauseRelation, nil
}

type SourceMechanismRelation struct {
	SourceID    int64 `db:"source_id"`
	MechanismID int64 `db:"mechanism_id"`
}

func CreateSourceMechanismRelation(sourceMechanismRelation *SourceMechanismRelation, db DB) error {
	q := `INSERT INTO source_mechanism_relation (source_id, mechanism_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceMechanismRelation.SourceID, sourceMechanismRelation.MechanismID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceMechanismRelationBySourceId(id, db DB) (*SourceMechanismRelation, error) {
	q := `SELECT * FROM source_mechanism_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceMechanismRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceMechanismRelation])
	if err != nil {
		return nil, err
	}

	return &sourceMechanismRelation, nil
}

func GetSourceMechanismRelationByMechanismId(id, db DB) (*SourceMechanismRelation, error) {
	q := `SELECT * FROM source_mechanism_relation WHERE mechanism_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceMechanismRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceMechanismRelation])
	if err != nil {
		return nil, err
	}

	return &sourceMechanismRelation, nil
}

type SourceSymptomRelation struct {
	SourceID  int64 `db:"source_id"`
	SymptomID int64 `db:"symptom_id"`
}

func CreateSourceSymptomRelation(sourceSymptomRelation *SourceSymptomRelation, db DB) error {
	q := `INSERT INTO source_symptom_relation (source_id, symptom_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceSymptomRelation.SourceID, sourceSymptomRelation.SymptomID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceSymptomRelationBySourceID(id, db DB) (*SourceSymptomRelation, error) {
	q := `SELECT * FROM source_symptom_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceSymptomRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceSymptomRelation])
	if err != nil {
		return nil, err
	}

	return &sourceSymptomRelation, nil
}

func GetSourceSymptomRelationBySymptomID(id, db DB) (*SourceSymptomRelation, error) {
	q := `SELECT * FROM source_symptom_relation WHERE symptom_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceSymptomRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceSymptomRelation])
	if err != nil {
		return nil, err
	}

	return &sourceSymptomRelation, nil
}

type SourceSymptomxRelation struct {
	SourceID   int64 `db:"source_id"`
	SymptomxID int64 `db:"symptomx_id"`
}

func CreateSourceSymptomxRelation(sourceSymptomRelation *SourceSymptomxRelation, db DB) error {
	q := `INSERT INTO source_symptomx_relation (source_id, symptomx_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceSymptomRelation.SourceID, sourceSymptomRelation.SymptomxID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceSymptomxRelationBySourceID(id, db DB) (*SourceSymptomxRelation, error) {
	q := `SELECT * FROM source_symptomx_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceSymptomxRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceSymptomxRelation])
	if err != nil {
		return nil, err
	}

	return &sourceSymptomxRelation, nil
}

func GetSourceSymptomxRelationBySymptomxID(id, db DB) (*SourceSymptomxRelation, error) {
	q := `SELECT * FROM source_symptomx_relation WHERE symptomx_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceSymptomxRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceSymptomxRelation])
	if err != nil {
		return nil, err
	}

	return &sourceSymptomxRelation, nil
}

type SourceMethodRelation struct {
	SourceID int64 `db:"source_id"`
	MethodID int64 `db:"method_id"`
}

func CreateSourceMethodRelation(sourceMethodRelation *SourceMethodRelation, db DB) error {
	q := `INSERT INTO source_method_relation (source_id, method_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceMethodRelation.SourceID, sourceMethodRelation.MethodID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceMethodRelationBySourceID(id, db DB) (*SourceMethodRelation, error) {
	q := `SELECT * FROM source_method_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceMethodRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceMethodRelation])
	if err != nil {
		return nil, err
	}

	return &sourceMethodRelation, nil
}

func GetSourceMethodRelationByMethodID(id, db DB) (*SourceMethodRelation, error) {
	q := `SELECT * FROM source_method_relation WHERE method_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceMethodRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceMethodRelation])
	if err != nil {
		return nil, err
	}

	return &sourceMethodRelation, nil
}

type SourceFormulaRelation struct {
	SourceID  int64 `db:"source_id"`
	FormulaID int64 `db:"formula_id"`
}

func CreateSourceFormulaRelation(sourceFormulaRelation *SourceFormulaRelation, db DB) error {
	q := `INSERT INTO source_formula_relation (source_id, formula_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceFormulaRelation.SourceID, sourceFormulaRelation.FormulaID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceFormulaRelationBySourceID(id, db DB) (*SourceFormulaRelation, error) {
	q := `SELECT * FROM source_formula_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceFormulaRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceFormulaRelation])
	if err != nil {
		return nil, err
	}

	return &sourceFormulaRelation, nil
}

func GetSourceFormulaRelationByFormulaID(id, db DB) (*SourceFormulaRelation, error) {
	q := `SELECT * FROM source_formula_relation WHERE formula_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceFormulaRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceFormulaRelation])
	if err != nil {
		return nil, err
	}

	return &sourceFormulaRelation, nil
}

type SourceHerbRelation struct {
	SourceID int64 `db:"source_id"`
	HerbID   int64 `db:"herb_id"`
}

func CreateSourceHerbRelation(sourceHerbRelation *SourceHerbRelation, db DB) error {
	q := `INSERT INTO source_herb_relation (source_id, herb_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceHerbRelation.SourceID, sourceHerbRelation.HerbID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceHerbRelationBySourceID(id, db DB) (*SourceHerbRelation, error) {
	q := `SELECT * FROM source_herb_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceHerbRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceHerbRelation])
	if err != nil {
		return nil, err
	}

	return &sourceHerbRelation, nil
}

func GetSourceHerbRelationByHerbID(id, db DB) (*SourceHerbRelation, error) {
	q := `SELECT * FROM source_herb_relation WHERE herb_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceHerbRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceHerbRelation])
	if err != nil {
		return nil, err
	}

	return &sourceHerbRelation, nil
}

type SourceUsageRelation struct {
	SourceID int64 `db:"source_id"`
	UsageID  int64 `db:"usage_id"`
}

func CreateSourceUsageRelation(sourceUsageRelation *SourceUsageRelation, db DB) error {
	q := `INSERT INTO source_usage_relation (source_id, usage_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceUsageRelation.SourceID, sourceUsageRelation.UsageID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceUsageRelationBySourceID(id, db DB) (*SourceUsageRelation, error) {
	q := `SELECT * FROM source_usage_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceUsageRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceUsageRelation])
	if err != nil {
		return nil, err
	}

	return &sourceUsageRelation, nil
}

func GetSourceUsageRelationByUsageID(id, db DB) (*SourceUsageRelation, error) {
	q := `SELECT * FROM source_usage_relation WHERE usage_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceUsageRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceUsageRelation])
	if err != nil {
		return nil, err
	}

	return &sourceUsageRelation, nil
}

type SourceVariationRelation struct {
	SourceID    int64 `db:"source_id"`
	VariationID int64 `db:"variation_id"`
}

func CreateSourceVariationRelation(sourceVariationRelation *SourceVariationRelation, db DB) error {
	q := `INSERT INTO source_variation_relation (source_id, variation_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceVariationRelation.SourceID, sourceVariationRelation.VariationID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceVariationRelationBySourceID(id, db DB) (*SourceVariationRelation, error) {
	q := `SELECT * FROM source_variation_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceVariationRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceVariationRelation])
	if err != nil {
		return nil, err
	}

	return &sourceVariationRelation, nil
}

func GetSourceVariationRelationByVariationID(id, db DB) (*SourceVariationRelation, error) {
	q := `SELECT * FROM source_variation_relation WHERE variation_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceVariationRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceVariationRelation])
	if err != nil {
		return nil, err
	}

	return &sourceVariationRelation, nil
}

type SourceContraindicationRelation struct {
	SourceID           int64 `db:"source_id"`
	ContraindicationID int64 `db:"contraindication_id"`
}

func CreateSourceContraindicationRelation(sourceContraindicationRelation *SourceContraindicationRelation, db DB) error {
	q := `INSERT INTO source_contraindication_relation (source_id, contraindication_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceContraindicationRelation.SourceID, sourceContraindicationRelation.ContraindicationID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceContraindicationRelationBySourceID(id, db DB) (*SourceContraindicationRelation, error) {
	q := `SELECT * FROM source_contraindication_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceContraindicationRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceContraindicationRelation])
	if err != nil {
		return nil, err
	}

	return &sourceContraindicationRelation, nil
}

func GetSourceContraindicationRelationByContraindicationID(id, db DB) (*SourceContraindicationRelation, error) {
	q := `SELECT * FROM source_contraindication_relation WHERE contraindication_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceContraindicationRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceContraindicationRelation])
	if err != nil {
		return nil, err
	}

	return &sourceContraindicationRelation, nil
}

type SourceAlternativeRelation struct {
	SourceID      int64 `db:"source_id"`
	AlternativeID int64 `db:"alternative_id"`
}

func CreateSourceAlternativeRelation(sourceAlternativeRelation *SourceAlternativeRelation, db DB) error {
	q := `INSERT INTO source_alternative_relation (source_id, alternative_id)
VALUES ($1, $2)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := db.Exec(ctx, q, sourceAlternativeRelation.SourceID, sourceAlternativeRelation.AlternativeID)

	if err != nil {
		return err
	}

	return nil
}

func GetSourceAlternativeRelationBySourceID(id, db DB) (*SourceAlternativeRelation, error) {
	q := `SELECT * FROM source_alternative_relation WHERE source_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceAlternativeRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceAlternativeRelation])
	if err != nil {
		return nil, err
	}

	return &sourceAlternativeRelation, nil
}

func GetSourceAlternativeRelationByAlternativeID(id, db DB) (*SourceAlternativeRelation, error) {
	q := `SELECT * FROM source_alternative_relation WHERE alternative_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := db.Query(ctx, q, id)
	if err != nil {
		return nil, err
	}

	sourceAlternativeRelation, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[SourceAlternativeRelation])
	if err != nil {
		return nil, err
	}

	return &sourceAlternativeRelation, nil
}
