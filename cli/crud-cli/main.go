package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pobyzaarif/belajar-go-cli/config"
)

func main() {
	mysqlDSN := os.Getenv("MYSQL_DSN")
	db := config.InitDatabase(mysqlDSN)
	defer db.Close()

}
