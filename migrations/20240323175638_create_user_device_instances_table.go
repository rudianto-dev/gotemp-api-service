package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserDeviceInstancesTable, downCreateUserDeviceInstancesTable)
}

func upCreateUserDeviceInstancesTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS user_device_instances
		(
			id VARCHAR (100) NOT NULL,
			user_id VARCHAR (100) NOT NULL,
			device_id VARCHAR (255) NOT NULL,
			instance_id VARCHAR (255) NOT NULL,
			created_at BIGINT NOT NULL DEFAULT 0,
			updated_at BIGINT NOT NULL DEFAULT 0,
			
			PRIMARY KEY (id),
			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateUserDeviceInstancesTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS user_device_instances;")
	if err != nil {
		return err
	}
	return nil
}
