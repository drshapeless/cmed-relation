-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS causes (
       id bigserial PRIMARY KEY,
       content text unique NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS causes;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
