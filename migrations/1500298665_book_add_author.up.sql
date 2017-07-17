ALTER TABLE book
    ADD COLUMN author_id INTEGER;
ALTER TABLE book
   ADD CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES person(id);
