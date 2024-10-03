package cmd

import (
	"fmt"

	"github.com/ahmedmohamed24/ecommerce-microservices/admin/config"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/database"
	"github.com/ahmedmohamed24/ecommerce-microservices/admin/internals/types"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// createAdminCmd represents the createAdmin command
var createAdminCmd = &cobra.Command{
	Use:   "createAdmin",
	Short: "create an admin user",
	Long:  `createAdmin is a command to create an admin user.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := CreateAdmin(cmd)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(createAdminCmd)
	createAdminCmd.Flags().StringP("name", "n", "", "Name of the admin")
	createAdminCmd.Flags().StringP("email", "e", "", "Email of the admin")
	createAdminCmd.Flags().StringP("password", "p", "", "Password of the admin")
}

type CreateAdminRequest struct {
	Name     string `validate:"required,min=3,max=100"`
	Email    string `validate:"required,email,min=3,max=100"`
	Password string `validate:"required,min=6,max=100"`
}

func CreateAdmin(cmd *cobra.Command) error {
	admin := &CreateAdminRequest{
		Name:     cmd.Flag("name").Value.String(),
		Email:    cmd.Flag("email").Value.String(),
		Password: cmd.Flag("password").Value.String(),
	}
	validate := validator.New()
	err := validate.Struct(admin)
	if err != nil {
		return err
	}
	config, err := config.New()
	if err != nil {
		return err
	}
	db, err := database.New(config)
	if err != nil {
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	if err := db.Create(&types.Admin{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: string(password),
	}).Error; err != nil {
		return err
	}
	fmt.Println("Admin created successfully...")
	return nil

}
