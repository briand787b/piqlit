DROP INDEX IF EXISTS parent_child_media_index;
DROP TABLE IF EXISTS parent_child_media;
DROP TABLE IF EXISTS media;

CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    length bigint NOT NULL,
    encoding VARCHAR(55) NOT NULL,
    upload_status VARCHAR(55) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX media_name 
ON media (name);

CREATE TABLE parent_child_media (
    parent_id INTEGER REFERENCES media(id) NOT NULL,
    child_id INTEGER REFERENCES media(id) NOT NULL
);

CREATE UNIQUE INDEX parent_child_media_index 
ON parent_child_media (parent_id, child_id);