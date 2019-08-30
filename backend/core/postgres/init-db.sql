DROP TABLE IF EXISTS servers;

CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(15) NOT NULL UNIQUE,
    is_master BOOLEAN NOT NULL DEFAULT false
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS media;

CREATE TABLE media (
    id SERIAL PRIMARY KEY,
    title VARCHAR(55) NOT NULL UNIQUE,
    thumbnail_name VARCHAR(55) NOT NULL,
    release_date TIMESTAMP NOT NULL,
    parent_id INTEGER REFERENCES media(id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- DROP TABLE IF EXISTS media_media;

-- CREATE TABLE media_media (
--     parent_id INTEGER REFERENCES media(id) NOT NULL,
--     child_id INTEGER REFERENCES media(id) NOT NULL
--     created_at TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

DROP TABLE IF EXISTS media_servers;

CREATE TABLE media_servers (
    server_id INTEGER REFERENCES servers(id) NOT NULL,
    media_id INTEGER REFERENCES media(id) NOT NULL,
);