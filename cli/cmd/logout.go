package cmd

import (
	"fmt"

	"github.com/Caroline1997/Service-Agenda/cli/service"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout current user",
	Long:  `the usage of the command is to log out current user`,
	Run: func(cmd *cobra.Command, args []string) {
		err := service.Logout()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Logout successfully!")
		}

	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
	logoutCmd.Flags().StringP("name", "n", "", "logout username")
	logoutCmd.Flags().StringP("password", "p", "", "password")
}
