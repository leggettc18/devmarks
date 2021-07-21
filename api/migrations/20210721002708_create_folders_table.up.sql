CREATE TABLE IF NOT EXISTS folders(
    id serial PRIMARY KEY,
    name text NOT NULL,
    color text,
    owner_id int NOT NULL,
    parent_id int,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    CONSTRAINT folders_owner_id_fkey FOREIGN KEY (owner_id)
    REFERENCES users(id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE CASCADE
);