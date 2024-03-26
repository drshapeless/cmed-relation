-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS sources (
       id bigserial PRIMARY KEY,
       content text NOT NULL,
       remark text NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS sources;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
