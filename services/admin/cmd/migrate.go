package cmd

import (
	"database/sql"
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate is a command to migrate the database",
	Long:  `migrate is a command to migrate the database. It can be used to create the tables in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		action, _ := cmd.Flags().GetString("action")
		step, _ := cmd.Flags().GetInt("step")
		err := Migrate(action, step)
		if err != nil {
			fmt.Println(err)
			return
		}
		if action == "up" {
			fmt.Println("Migration successful")
		} else {
			fmt.Println("Rollback successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringP("action", "a", "up", "Migration action (up/down)")
	migrateCmd.Flags().IntP("step", "s", 0, "Migration step to execute or rollback")
}

func Migrate(action string, step int) error {
	config, err := config.New()
	if err != nil {
		return err
	}
	host := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName)
	db, err := sql.Open("postgres", host)
	if err != nil {
		return err
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: config.Database.DBName,
	})
	if err != nil {
		return err
	}
	defer driver.Close()
	m, err := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)
	if err != nil {
		return err
	}
	if action == "down" {
		if step == 0 {
			return m.Down()
		}
		return m.Steps(-step)
	} else {
		if step == 0 {
			return m.Up()
		}
		return m.Steps(step)
	}
}
