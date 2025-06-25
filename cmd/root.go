package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "td",
	Short: "Shortcuts for frequently used commands",
	Long: `td provides shortcuts for frequently used commands. 
	       Source code available at https://github.com/thisistonydang/td`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
