package main

import "grafana-manager/pkg/logger"

func main() {

	logger.Init("debug")
	Execute()
}
