-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS source_disease_relation (
       source_id bigint NOT NULL,
       disease_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_cause_relation (
       source_id bigint NOT NULL,
       cause_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_mechanism_relation (
       source_id bigint NOT NULL,
       mechanism_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_symptom_relation (
       source_id bigint NOT NULL,
       symptom_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_symptomx_relation (
       source_id bigint NOT NULL,
       symptomx_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_method_relation (
       source_id bigint NOT NULL,
       method_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_formula_relation (
       source_id bigint NOT NULL,
       formula_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_herb_relation (
       source_id bigint NOT NULL,
       herb_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_usage_relation (
       source_id bigint NOT NULL,
       usage_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_variation_relation (
       source_id bigint NOT NULL,
       variation_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_contraindication_relation (
       source_id bigint NOT NULL,
       contraindication_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS source_alternative_relation (
       source_id bigint NOT NULL,
       alternative_id bigint NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS source_disease_relation;
DROP TABLE IF EXISTS source_cause_relation;
DROP TABLE IF EXISTS source_mechanism_relation;
DROP TABLE IF EXISTS source_symptom_relation;
DROP TABLE IF EXISTS source_symptomx_relation;
DROP TABLE IF EXISTS source_method_relation;
DROP TABLE IF EXISTS source_formula_relation;
DROP TABLE IF EXISTS source_herb_relation;
DROP TABLE IF EXISTS source_usage_relation;
DROP TABLE IF EXISTS source_variation_relation;
DROP TABLE IF EXISTS source_contraindication_relation;
DROP TABLE IF EXISTS source_alternative_relation;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
