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
			id VARCHAR (100) NOT NULL,
			name VARCHAR (255) NOT NULL,
			username VARCHAR (100) NOT NULL,
			password VARCHAR (100) NOT NULL DEFAULT '',
			status SMALLINT NOT NULL DEFAULT 0,
			role_type SMALLINT NOT NULL DEFAULT 0,
			created_at BIGINT NOT NULL DEFAULT 0,
			updated_at BIGINT NOT NULL DEFAULT 0,
			deleted_at BIGINT NOT NULL DEFAULT 0,

			PRIMARY KEY (id),
			UNIQUE (username)
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
