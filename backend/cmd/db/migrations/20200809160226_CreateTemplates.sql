-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE templates (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title TEXT NOT NULL,
   content TEXT DEFAULT NULL,
   created_at TEXT NOT NULL,
   updated_at TEXT NOT NULL
);
CREATE INDEX template_id on templates (id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX template_id;
DROP TABLE templates;
