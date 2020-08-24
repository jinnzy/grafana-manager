package manager

import (
	"grafana-manager/cmd/manager/options"
	"log"
	"os"
)

type Manager struct {
	URL  string
	BasicAuthOrToken string
}

func NewManager(option *options.Options) *Manager {
	var manager Manager

	if  option.Token == "" && option.User == ""  && option.Password == ""  {
		log.Fatal("--token or --username and --password cannot be empty")
		os.Exit(1)
	}
	if option.Token != "" {
		manager.BasicAuthOrToken = option.Token
	}
	if option.User != ""  && option.Password != "" {
		manager.BasicAuthOrToken = option.User + ":" + option.Password
	}
	manager.URL = option.URL

	return &manager
}


//func (m *Manager)Validation() error {
//
//}
