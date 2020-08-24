package main

import (
	"github.com/spf13/cobra"
	"grafana-manager/pkg/manager"
)

func init()  {
	importCmd.AddCommand(datasourceCmd)
}


var datasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "import datasource",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gManager := manager.NewManager(&mOptions)
		gManager.ImportDatasource(args[0])
	},
}
