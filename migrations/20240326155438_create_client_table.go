package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateClientTable, downCreateClientTable)
}

func upCreateClientTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS clients
		(
			id VARCHAR (100) NOT NULL,
			client_id VARCHAR (100) NOT NULL,
			client_secret VARCHAR (100) NOT NULL,
			status SMALLINT NOT NULL DEFAULT 0,
			expired_at BIGINT NOT NULL DEFAULT 0,
			
			PRIMARY KEY (id),
			UNIQUE (client_id)
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateClientTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS clients;")
	if err != nil {
		return err
	}
	return nil
}
