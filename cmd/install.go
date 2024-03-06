/*
Copyright © 2024 Jason Dsouza <jason@leanj.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs this utility",
	Long: `Installs this utility
The following lines depict usage of this utillity:

ljd install`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("installing")
		installationDir, err := os.UserHomeDir()
		if err == nil {
			fmt.Print(installationDir)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
