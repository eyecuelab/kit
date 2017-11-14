package migrate

import (
	"fmt"
	"os"

	"github.com/eyecuelab/kit/assets"
	"github.com/eyecuelab/kit/db/psql"
	"github.com/eyecuelab/kit/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

const migrationsDir = "data/bin/migrations"

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations on the database",
	Long:  `Can be used to forward and rollback migrations.`,
	Run:   run,
}

var (
	down, files bool
	max         int
)

func init() {
	MigrateCmd.Flags().BoolVar(&down, "down", false, "rollback migrations")
	MigrateCmd.Flags().IntVar(&max, "max", 0, "number of migrations to run, default will run all")
	MigrateCmd.Flags().BoolVar(&files, "files", false, fmt.Sprintf("use files directly rather than bindata"))
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 1 && args[0] == "config" {
		writeConfig()
		return
	}

	dir := direction()
	n, err := migrate.ExecMax(psql.DB.DB(), "postgres", getMigrations(), dir, max)
	if err != nil {
		log.Fatalf("migrate.run: migrate.ExecMax: %v", err)
	}
	if down {
		log.Infof("Rolled-back %d migrations.\n", n)
	} else {
		log.Infof("Applied %d migrations.\n", n)
	}
}

func MigrateAll() (int, error) {
	return migrate.Exec(psql.DB.DB(), "postgres", getMigrations(), migrate.Up)
}

func PendingMigrationCount() (int, error) {
	plans, _, err := migrate.PlanMigration(psql.DB.DB(), "postgres", getMigrations(), migrate.Up, 0)
	if err != nil {
		return -1, err
	}
	return len(plans), nil
}

func direction() migrate.MigrationDirection {
	if down {
		if max == 0 {
			max = 1
		}
		return migrate.Down
	}
	return migrate.Up
}

func migrationsDirExists() bool {
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return false
	}
	return true
}

func getMigrations() migrate.MigrationSource {
	if files && migrationsDirExists() {
		return &migrate.FileMigrationSource{Dir: migrationsDir}
	}
	return &migrate.AssetMigrationSource{
		Asset:    assets.Manager.Get,
		AssetDir: assets.Manager.Dir,
		Dir:      migrationsDir,
	}
}
