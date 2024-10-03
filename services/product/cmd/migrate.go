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

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run database migrations for products service",
	Long:  `This command will run database migrations for products service.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := configs.New()
		if err != nil {
			panic(err)
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", c.Database.Host, c.Database.Username, c.Database.Password, c.Database.Database, c.Database.Port)
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			fmt.Println("failed to connect to database: " + err.Error())
			return
		}
		driver, err := postgres.WithInstance(db, &postgres.Config{
			DatabaseName: c.Database.Database,
		})
		if err != nil {
			panic("failed to create driver: " + err.Error())
		}
		m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "postgres", driver)
		if err != nil {
			fmt.Println("failed to create migration: " + err.Error())
			return
		}

		step, err := strconv.Atoi(cmd.Flag("step").Value.String())
		if err != nil {
			fmt.Println("invalid step value: " + err.Error())
			return
		}
		if step > 0 {
			err = m.Steps(step)
			if err != nil {
				fmt.Println("failed to run migration steps: " + err.Error())
				return
			}
		}
		err = m.Up()
		if err != nil {
			fmt.Println("failed to run migration: " + err.Error())
			return
		}
		fmt.Println("migrate completed ....")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Int("step", 0, "migration steps")
}
