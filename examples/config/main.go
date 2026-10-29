package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rafabcanedo/obsidian-golang-toolkit/config"
)

type AppConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	envFile, _ := os.CreateTemp("", "*.env")
	envFile.WriteString("DB_HOST=localhost\nDB_PORT=5432\n")
	envFile.Close()
	defer os.Remove(envFile.Name())

	vars, err := config.LoadEnv(envFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB_HOST:", vars["DB_HOST"])
	fmt.Println("DB_PORT:", vars["DB_PORT"])

	jsonFile, _ := os.CreateTemp("", "*.json")
	jsonFile.WriteString(`{"host":"localhost","port":8080}`)
	jsonFile.Close()
	defer os.Remove(jsonFile.Name())

	var cfg AppConfig
	if err := config.LoadJSON(jsonFile.Name(), &cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("server: %s:%d\n", cfg.Host, cfg.Port)
}
