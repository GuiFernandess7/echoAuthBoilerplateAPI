// cmd/api/main.go
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
		fmt.Fprintf(os.Stderr, "Erro crítico ao carregar configurações: %v\n", err)
		os.Exit(1)
	}

	application, err := app.NewApplication(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro crítico ao inicializar a aplicação: %v\n", err)
		os.Exit(1)
	}

	defer application.Close()
	application.Run()
}
