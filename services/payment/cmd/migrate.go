package cmd

import (
	"fmt"
	"log"

	"github.com/ahmedmohamed24/ecommerce-microservices/payment/config"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Order Action",
	Long:  `Handle order operations like creating an order,`,
	Run: func(cmd *cobra.Command, args []string) {
		action, err := cmd.Flags().GetString("action")
		if err != nil || (action != "up" && action != "down") {
			log.Println("action should be up/down")
		}
		steps, err := cmd.Flags().GetInt8("step")
		if err != nil {
			log.Println("steps flag is not provided")
		}
		c, err := config.New()
		if err != nil {
			log.Println(err)
			return
		}
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Database)
		m, err := migrate.New("file://database/migrations", dsn)
		if err != nil {
			log.Println(err)
			return
		}
		if steps != 0 {
			err = m.Steps(int(steps))
			if err != nil {
				log.Println(err)
			} else {
				log.Println("migrated successfully")
			}
			return
		}
		if action == "down" {
			err := m.Down()
			if err != nil {
				log.Println(err)
			} else {
				log.Println("rollback done successfully")
			}
			return
		}
		err = m.Up()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("migrated successfully")

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Int8("step", 0, "steps to migrate the database, default is 0 and could be any number")
	migrateCmd.Flags().String("action", "up", "action to take on the database, default is up and could be up or down")
}
