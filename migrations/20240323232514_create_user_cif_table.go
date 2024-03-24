package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserCifTable, downCreateUserCifTable)
}

func upCreateUserCifTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS user_cif
		(
			id VARCHAR (100) NOT NULL,
			user_id VARCHAR (100) NOT NULL,
			reference_id VARCHAR (100) NOT NULL,
			created_at BIGINT NOT NULL DEFAULT 0,
			updated_at BIGINT NOT NULL DEFAULT 0,
			
			PRIMARY KEY (id),
			UNIQUE (reference_id),
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUserCifTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS user_cif;")
	if err != nil {
		return err
	}
	return nil
}
