package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ahmedmohamed24/ecommerce-microservices/customer/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migration command",
	Long:  `This command is used to migrate the database schema`,
	Run: func(cmd *cobra.Command, args []string) {
		action, _ := cmd.Flags().GetString("action")
		step, _ := cmd.Flags().GetInt("step")
		if step < 0 {
			log.Println("step must be greater than 0")
			return
		}
		if action != "up" && action != "down" {
			log.Println("action must be up or down")
			return
		}
		run(action, step)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringP("action", "a", "up", "migration action")
	migrateCmd.Flags().IntP("step", "s", 0, "migration step")
}

func run(action string, step int) {
	config, err := config.New()
	if err != nil {
		log.Println("error reading config: ", err)
		return
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	db, err := sql.Open("postgres", dsn)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Println("error connecting to database: ", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)
	if err != nil {
		log.Println("error connecting to database: ", err)
		return
	}
	if action == "up" {
		if step > 0 {
			err = m.Steps(step)
		} else {
			err = m.Up()
		}
	} else {
		if step > 0 {
			err = m.Steps(-step)
		} else {
			err = m.Down()
		}

	}

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no change")
			return
		}
		if action == "up" {
			log.Println("error while migrating up: ", err)
		} else {
			log.Println("error while migrating down: ", err)
		}
		return
	}
	if action == "up" {
		log.Println("migrated up successfully")
	} else {
		log.Println("migrated down successfully")
	}
}
