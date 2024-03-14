package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserTable, downCreateUserTable)
}

func upCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS users
		(
			id CHARACTER VARYING(255) NOT NULL,
			name CHARACTER VARYING(255) NOT NULL,
			created_at BIGINT NOT NULL DEFAULT 0,
			updated_at BIGINT NOT NULL DEFAULT 0,

			PRIMARY KEY (id)
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		return err
	}
	return nil
}
