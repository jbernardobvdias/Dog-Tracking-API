package main

import (
	"dog-tracking/data"
	"log"
)

func main() {
	log.Println("Setting up SQLite DB...")

	data.CreateTable()

	log.Println("Table set up.")
	log.Println("Setting up API...")

	data.SetupHttp("8080")
	// or
	// data.SetupGin("8080")

	log.Println("API set up.")
}
