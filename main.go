package main

import (
	"dog-tracking/data"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	log.Println("Setting up SQLite DB...")

	data.CreateTable()

	log.Println("Table set up.")
	log.Println("Setting up API...")

	data.SetupHttp(port)
	// or
	// data.SetupGin(port)

	log.Println("API set up.")
}
