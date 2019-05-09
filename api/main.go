package main

import (
	"github.com/eyalkenig/meizam-api/api/app/cmd"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	cmd.Execute(port)
}
