package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up001, Down001)
}

// Up001 applies this migration
func Up001(tx *sql.Tx) error {
	_, err := tx.Exec(`
    CREATE TABLE templates (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT DEFAULT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	 );
	 CREATE INDEX template_id on templates (id);
  `)
	return err
}

// Down001 downgrades this migration
func Down001(tx *sql.Tx) error {
	_, err := tx.Exec(`
	DROP INDEX template_id;
	DROP TABLE templates;
	`)
	return err
}