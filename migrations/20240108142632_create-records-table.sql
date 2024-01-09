-- migrate:up
CREATE TABLE IF NOT EXISTS records (
    id bigserial NOT NULL,
    name text NOT NULL,
    marks int [] NOT NULL,
    created_at timestamptz NOT NULL DEFAULT current_timestamp,
    PRIMARY KEY (id)
);
-- seed table
INSERT INTO records (name, marks) VALUES ('Gavinda', ARRAY[100, 100]);
INSERT INTO records (name, marks) VALUES ('Kinandana', ARRAY[50, 100]);
INSERT INTO records (name, marks) VALUES ('Gavinda Kinandana', ARRAY[100, 200, 300]);
-- add index to speed up query
CREATE INDEX records_marks_idx ON records (marks,created_at);

-- migrate:down
DROP INDEX records_marks_idx;
DROP TABLE IF EXISTS records;
