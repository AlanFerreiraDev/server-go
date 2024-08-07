package main

import (
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	migrationsPath := "/home/alan/Documentos/Projects/Go+React/server-go/internal/store/pgstore/migrations"
	configPath := "/home/alan/Documentos/Projects/Go+React/server-go/internal/store/pgstore/migrations/tern.conf"

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		migrationsPath,
		"--config",
		configPath,
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
