package infrastructure

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	_ "github.com/rudianto-dev/gotemp-api-service/migrations"
)

func (infra *Service) MigrateDB() error {
	infra.Logger.Info("running migration up")
	db, err := sql.Open(infra.DB.Config.DriverName, infra.DB.Config.SourceMaster)
	if err != nil {
		infra.Logger.Errorf("error connecting to db when running migration up, (%v)", err)
		return err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		infra.Logger.Errorf("error running migration up, (%v)", err)
	}
	infra.Logger.Info("running migration successfully")
	return nil
}

func (infra *Service) RollbackDB() error {
	infra.Logger.Info("running migration down")
	db, err := sql.Open(infra.DB.Config.DriverName, infra.DB.Config.SourceMaster)
	if err != nil {
		infra.Logger.Errorf("error connecting to db when running migration down, (%v)", err)
		return err
	}
	if err := goose.Down(db, "migrations"); err != nil {
		infra.Logger.Errorf("error running migration down, (%v)", err)
	}
	infra.Logger.Info("rollback migration successfully")
	return nil
}
