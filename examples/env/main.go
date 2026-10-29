package main

import (
	"fmt"
	"os"

	"github.com/rafabcanedo/obsidian-golang-toolkit/env"
)

func main() {
	host := env.GetString("DB_HOST", "localhost")
	port := env.GetString("DB_PORT", "5432")
	fmt.Printf("connecting to %s:%s\n", host, port)

	os.Setenv("APP_ENV", "production")
	appEnv := env.GetString("APP_ENV", "development")
	fmt.Println("environment:", appEnv)
}
