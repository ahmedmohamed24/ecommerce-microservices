package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/ahmedmohamed24/ecommerce-microservices/product/configs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:   "migrateDown",
	Short: "migrate down the database",
	Long:  `This command will migrate down the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := configs.New()
		if err != nil {
			fmt.Println("Failed to load configurations: ", err)
			return
		}
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.Database)
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("Failed to open connection: ", err)
			return
		}
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			fmt.Println("Failed to create driver: ", err)
			return
		}
		m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
		if err != nil {
			fmt.Println("Failed to create migration instance: ", err)
			return
		}
		steps, err := strconv.Atoi(cmd.Flag("step").Value.String())
		if err != nil {
			fmt.Println("Failed to parse step value: ", err)
			return
		}
		if steps == 0 {
			if err := m.Down(); err != nil {
				fmt.Println("Failed to migrate down: ", err)
				return
			}
		} else {
			if err := m.Steps(-steps); err != nil {
				fmt.Println("Failed to migrate down: ", err)
				return
			}

		}
		fmt.Println("Migrated down successfully...")
	},
}

func init() {
	rootCmd.AddCommand(migrateDownCmd)
	migrateDownCmd.Flags().String("step", "0", "number of steps to migrate down")
}
