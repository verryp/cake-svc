package cmd

import (
	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"github.com/verryp/cake-store-service/app/common"
	"github.com/verryp/cake-store-service/app/driver"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrateup",
	Short: "Migrate Up DB",
	Long:  "Run all of your outstanding migrations",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := common.NewConfig()
		if err != nil {
			log.Logger.Err(err).Msgf("Config error | %v", err)
			panic(err)
		}

		src := getMigrateSrc()

		doMigrate(conf.DB, src, migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Down DB",
	Long:  "Rollback all the migrations",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := common.NewConfig()
		if err != nil {
			log.Logger.Err(err).Msgf("Config error | %v", err)
			panic(err)
		}

		src := getMigrateSrc()

		doMigrate(conf.DB, src, migrate.Down)
	},
}

func getMigrateSrc() migrate.FileMigrationSource {
	src := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return src
}

func doMigrate(conf common.DB, mSource migrate.FileMigrationSource, direction migrate.MigrationDirection) error {
	db, err := driver.NewMySQLDatabase(conf)
	if err != nil {
		log.Err(err).Msg("error on db connection")
		return err
	}

	defer db.Db.Close()

	total, err := migrate.Exec(db.Db, "mysql", mSource, direction)
	if err != nil {
		log.Err(err).Msg("error when do migration")
		return err
	}

	log.Info().Msgf("Migrate success, total migrated: %d", total)
	return nil
}

func init() {
	RootCmd.AddCommand(migrateUpCmd)
	RootCmd.AddCommand(migrateDownCmd)
}
