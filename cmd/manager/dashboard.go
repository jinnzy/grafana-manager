package main

import (
	"github.com/spf13/cobra"
	"grafana-manager/pkg/manager"
)

func init()  {
	importCmd.AddCommand(dashboardCmd)
}

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "import dashboard",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gManager := manager.NewManager(&mOptions)
		gManager.ImportDashboard(args[0])
	},
}
