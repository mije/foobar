ALTER TABLE book
    DROP CONSTRAINT fk_author;
ALTER TABLE book
    DROP book COLUMN author_id;
