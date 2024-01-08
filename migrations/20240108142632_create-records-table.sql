-- migrate:up
CREATE TABLE IF NOT EXISTS public.records (
    id bigserial NOT NULL,
    name text NOT NULL,
    marks int [] NOT NULL,
    createdat timestamptz NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (id)
);

-- migrate:down
DROP TABLE IF EXISTS records;
