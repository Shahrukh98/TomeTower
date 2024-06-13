package main

import (
	"log"

	"tometower/internal/app"
)

func main() {
	app := app.NewApp()
	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
