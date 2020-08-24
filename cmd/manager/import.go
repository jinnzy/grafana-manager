package main

import (
	"github.com/spf13/cobra"
)


func init()  {
	rootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import dashboard or datasource",
}


