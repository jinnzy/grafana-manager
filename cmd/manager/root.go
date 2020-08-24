package main

import (
	"github.com/spf13/cobra"
	"grafana-manager/cmd/manager/options"
	"log"
	"os"
)

func init()  {
	rootCmd.PersistentFlags().StringVar(&mOptions.URL, "url", "", "http://grafana:3000")
	rootCmd.PersistentFlags().StringVar(&mOptions.Token,"token", "", "token")
	rootCmd.PersistentFlags().StringVarP(&mOptions.User,"user", "u","", "token")
	rootCmd.PersistentFlags().StringVarP(&mOptions.Password,"password", "p", "", "token")
	// 必须存在url
	if err := rootCmd.MarkPersistentFlagRequired("url");err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
var mOptions = options.Options{}

var rootCmd = &cobra.Command{
	Use:   "grafana-manager",
	Short: "grafana manager",
	Long: `grafana manager`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//log.Print(err)
		os.Exit(1)
	}
}
