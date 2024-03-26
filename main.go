package main

import (
	"bufio"
	"cmed-relation/data"
	"context"
	"encoding/csv"
	"errors"
	"os"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	print("connecting to database ")
	println(dbURL)
	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		println("cannot connect to database")
		os.Exit(1)
	}
	defer dbpool.Close()

	reader := bufio.NewReader(os.Stdin)

	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		println("csv fuck up")
	}

	for _, v := range records {
		// 0. source
		source := data.Source{
			Content: v[0],
			Remark:  v[13],
		}
		_ = data.CreateSource(&source, dbpool)

		// 1. disease
		_ = processDiseases(v[1], &source, dbpool)

		// 2. cause
		_ = processCauses(v[2], &source, dbpool)

		// 3. mechanism
		_ = processMechanisms(v[3], &source, dbpool)

		// 4. symptom
		_ = processSymptoms(v[4], &source, dbpool)

		// 5. symptomx
		_ = processSymptomx(v[5], &source, dbpool)

		// 6. method
		_ = processMethods(v[6], &source, dbpool)

		// 7. formula
		_ = processFormulae(v[7], &source, dbpool)

		// 8. herb
		_ = processHerbs(v[8], &source, dbpool)

		// 9. usage
		_ = processUsages(v[9], &source, dbpool)

		// 10. variation
		_ = processVariations(v[10], &source, dbpool)

		// 11. contraindication
		_ = processContraindications(v[11], &source, dbpool)

		// 12. alternative
		_ = processAlternatives(v[12], &source, dbpool)
	}
}

func processDiseases(s string, source *data.Source, db data.DB) error {
	diseases := strings.Split(s, "\n\n")
	for _, v := range diseases {
		if v != "" {
			d := data.Disease{
				Content: v,
			}
			err := data.CreateDisease(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetDiseaseByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd
				} else {
					return err
				}
			}

			r := data.SourceDiseaseRelation{
				SourceID:  source.ID,
				DiseaseID: d.ID,
			}
			err = data.CreateSourceDiseaseRelation(&r, db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func processCauses(s string, source *data.Source, db data.DB) error {
	causes := strings.Split(s, "\n\n")
	for _, v := range causes {
		if v != "" {
			d := data.Cause{
				Content: v,
			}
			err := data.CreateCause(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetCauseByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd
				} else {
					return err
				}
			}

			r := data.SourceCauseRelation{
				SourceID: source.ID,
				CauseID:  d.ID,
			}
			err = data.CreateSourceCauseRelation(&r, db)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func processMechanisms(s string, source *data.Source, db data.DB) error {
	mechanisms := strings.Split(s, "\n\n")
	for _, v := range mechanisms {
		if v != "" {
			d := data.Mechanism{
				Content: v,
			}

			err := data.CreateMechanism(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetMechanismByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd
				} else {
					return err
				}
			}

			r := data.SourceMechanismRelation{
				SourceID:    source.ID,
				MechanismID: d.ID,
			}
			err = data.CreateSourceMechanismRelation(&r, db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func processSymptoms(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Symptom{
				Content: v,
			}

			err := data.CreateSymptom(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetSymptomByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceSymptomRelation{
				SourceID:  source.ID,
				SymptomID: d.ID,
			}
			err = data.CreateSourceSymptomRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processSymptomx(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Symptomx{
				Content: v,
			}

			err := data.CreateSymptomx(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetSymptomxByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceSymptomxRelation{
				SourceID:   source.ID,
				SymptomxID: d.ID,
			}
			err = data.CreateSourceSymptomxRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processMethods(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Method{
				Content: v,
			}

			err := data.CreateMethod(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetMethodByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceMethodRelation{
				SourceID: source.ID,
				MethodID: d.ID,
			}
			err = data.CreateSourceMethodRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processFormulae(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Formulae{
				Content: v,
			}

			err := data.CreateFormulae(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetFormulaeByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceFormulaRelation{
				SourceID:  source.ID,
				FormulaID: d.ID,
			}
			err = data.CreateSourceFormulaRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processHerbs(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n")
	for _, v := range ss {
		if v != "" {
			// The syntax of herb is like this herbname(process)
			// The process is not needed

			if strings.Contains(v, "（") {
				tmp := strings.Split(v, "（")
				v = strings.TrimSpace(tmp[0])
			}

			d := data.Herb{
				Content: v,
			}

			err := data.CreateHerb(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetHerbByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceHerbRelation{
				SourceID: source.ID,
				HerbID:   d.ID,
			}
			err = data.CreateSourceHerbRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processUsages(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Usage{
				Content: v,
			}

			err := data.CreateUsage(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetUsageByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceUsageRelation{
				SourceID: source.ID,
				UsageID:  d.ID,
			}
			err = data.CreateSourceUsageRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processVariations(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Variation{
				Content: v,
			}

			err := data.CreateVariation(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetVariationByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceVariationRelation{
				SourceID:    source.ID,
				VariationID: d.ID,
			}
			err = data.CreateSourceVariationRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processContraindications(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Contraindication{
				Content: v,
			}

			err := data.CreateContraindication(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetContraindicationByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceContraindicationRelation{
				SourceID:           source.ID,
				ContraindicationID: d.ID,
			}
			err = data.CreateSourceContraindicationRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

func processAlternatives(s string, source *data.Source, db data.DB) error {
	ss := strings.Split(s, "\n\n")
	for _, v := range ss {
		if v != "" {
			d := data.Alternative{
				Content: v,
			}

			err := data.CreateAlternative(&d, db)
			if err != nil {
				var e *pgconn.PgError
				if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
					dd, err := data.GetAlternativeByContent(v, db)
					if err != nil {
						return err
					}
					d = *dd

				} else {
					return err
				}
			}

			r := data.SourceAlternativeRelation{
				SourceID:      source.ID,
				AlternativeID: d.ID,
			}
			err = data.CreateSourceAlternativeRelation(&r, db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
