package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserPhoneNumbersTable, downCreateUserPhoneNumbersTable)
}

func upCreateUserPhoneNumbersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS user_phone_numbers
		(
			id VARCHAR (100) NOT NULL,
			user_id VARCHAR (100) NOT NULL,
			phone_number VARCHAR (25) NOT NULL,
			created_at BIGINT NOT NULL DEFAULT 0,
			updated_at BIGINT NOT NULL DEFAULT 0,
			deleted_at BIGINT NOT NULL DEFAULT 0,
			
			PRIMARY KEY (id),
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			UNIQUE (phone_number)
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUserPhoneNumbersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS user_phone_numbers;")
	if err != nil {
		return err
	}
	return nil
}
