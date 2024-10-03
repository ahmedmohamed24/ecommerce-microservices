package cmd

import (
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/scheduler"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server. This command starts the server on the port specified in the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		scheduler.Init()
		err := server.Serve()
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
