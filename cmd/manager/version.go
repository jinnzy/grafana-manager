package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of grafana-manager",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grafana-manager Static Site Generator v0.0.0 -- HEAD")
	},
}
