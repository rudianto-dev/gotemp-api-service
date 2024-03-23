package main

import (
	"log"

	"github.com/rudianto-dev/gotemp-api-service/cmd/configuration"
	"github.com/rudianto-dev/gotemp-api-service/cmd/infrastructure"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{Use: "boilerplate", Short: "Go Boilerplate"}
	configuration := configuration.New()
	infrastructure := infrastructure.New(configuration)

	cmd.AddCommand(&cobra.Command{
		Use:   "api",
		Short: "Run API Service",
		Run: func(*cobra.Command, []string) {
			infrastructure.RunAPI()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "migrate:up",
		Short: "Run Migration",
		RunE: func(*cobra.Command, []string) error {
			return infrastructure.MigrateDB()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "migrate:down",
		Short: "Rollback Migration",
		RunE: func(*cobra.Command, []string) error {
			return infrastructure.RollbackDB()
		},
	})

	if err := cmd.Execute(); err != nil {
		log.Panic(err.Error())
	}
}
