CREATE TABLE IF NOT EXISTS bookmark_folder(
    bookmark_id int NOT NULL,
    folder_id int NOT NULL,

    CONSTRAINT bookmarks_id_fkey FOREIGN KEY (bookmark_id)
    REFERENCES bookmarks(id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE CASCADE,

    CONSTRAINT folders_id_fkey FOREIGN KEY (folder_id)
    REFERENCES folders(id) MATCH SIMPLE
    ON UPDATE NO ACTION ON DELETE CASCADE
);