-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS contraindications (
       id bigserial PRIMARY KEY,
       content text UNIQUE NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS contraindications;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
