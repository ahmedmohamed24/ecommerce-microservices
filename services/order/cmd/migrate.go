package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ahmedmohamed24/ecommerce-microservices/order/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "running migrations",
	Long:  `This command will run the migrations for the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.New()
		if err != nil {
			log.Println(err.Error())
			return
		}
		action, _ := cmd.Flags().GetString("action")
		step, _ := cmd.Flags().GetInt("step")

		dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
		db, err := sql.Open("postgres", dns)
		if err != nil {
			log.Println(err.Error())
			return
		}
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Println(err.Error())
			return
		}

		m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if step != 0 {
			err = m.Steps(int(step))
		} else {
			if action == "up" {
				err = m.Up()
			} else {
				err = m.Down()
			}
		}
		if err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println("Migration run successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().String("action", "up", "Action to run the migrations up/down")
	migrateCmd.Flags().Int("step", 0, "migration step to run should be greater than 0")
}
