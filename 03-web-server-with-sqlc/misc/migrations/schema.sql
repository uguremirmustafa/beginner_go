DROP TABLE IF EXISTS nodes;

CREATE TABLE IF NOT EXISTS nodes (
    id INTEGER PRIMARY KEY,
    name text NOT NULL,
    description text
);