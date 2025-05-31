package main

import (
	"fmt"
	"os"

	"github.com/GuiFernandess7/echoAuthBoilerplate/internal/app"
	"github.com/GuiFernandess7/echoAuthBoilerplate/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading application configuration: %v\n", err)
		os.Exit(1)
	}

	application, err := app.NewApplication(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
		os.Exit(1)
	}

	defer application.Close()
	application.Run()
}
