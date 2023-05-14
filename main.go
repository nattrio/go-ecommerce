package main

import (
	"fmt"
	"os"

	"github.com/nattrio/go-ecommerce/config"
)

func envPath() string {
	if len(os.Args) == 1 { // no args
		return ".env"
	} else {
		return os.Args[1] // first arg
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	fmt.Println(cfg.App())
	fmt.Println(cfg.Db())
	fmt.Println(cfg.Jwt())
}
