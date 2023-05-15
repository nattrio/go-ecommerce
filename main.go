package main

import (
	"fmt"
	"os"

	"github.com/nattrio/go-ecommerce/config"
	"github.com/nattrio/go-ecommerce/pkg/databases"
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

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	fmt.Println(db)
}
