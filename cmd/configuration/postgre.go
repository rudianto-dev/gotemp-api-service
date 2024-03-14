package configuration

import (
	"github.com/rudianto-dev/gotemp-sdk/pkg/database"
)

func (cf ConfigurationSchema) NewPostgres() database.DatabaseConfig {
	return database.DatabaseConfig{
		DriverName:              database.Postgres,
		SourceMaster:            cf.Database.Master,
		SourceSlave:             cf.Database.Slave,
		IntervalCheckConnection: cf.Database.IntervalCheckConnection,
	}
}
