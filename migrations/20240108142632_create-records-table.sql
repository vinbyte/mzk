-- migrate:up
CREATE TABLE IF NOT EXISTS records (
    id bigserial NOT NULL,
    name text NOT NULL,
    marks int [] NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (id)
);
INSERT INTO records (name, marks) VALUES ('Gavinda', ARRAY[100, 300, 400]);
INSERT INTO records (name, marks) VALUES ('Kinandana', ARRAY[200, 300, 500]);

-- migrate:down
DROP TABLE IF EXISTS records;
