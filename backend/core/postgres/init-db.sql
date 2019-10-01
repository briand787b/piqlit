DROP TABLE IF EXISTS media;

CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    encoding VARCHAR(55) NOT NULL,
    upload_status VARCHAR(55) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS parent_child_media;

CREATE TABLE parent_child_media (
    parent_id INTEGER REFERENCES media(id) NOT NULL,
    child_id INTEGER REFERENCES media(id) NOT NULL
);