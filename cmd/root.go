package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "IPtracker CLI application",
	Long:  `IPtracker CLI application, which will display location inforamtion given the IP address`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("IPtracker CLI application")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
