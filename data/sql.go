package data

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const DBTYPE string = "sqlite3"

var DBPATH string = os.Getenv("DBPATH")

func CreateTable() {
	// In case it is run locally and DBPATH doesn't exist
	if DBPATH == "" {
		DBPATH = "./database.db"
	}

	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	// Create the dogs table if it does not exist.
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS dogs (id INTEGER PRIMARY KEY, name TEXT UNIQUE, race TEXT, age INTEGER)")

	if err != nil {
		log.Fatal("There was an issue with setting up the SQL statement. " + err.Error())
	}
	statement.Exec()
}

func GetDogs() [][]string {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dogs")

	if err != nil {
		log.Fatal("There was a problem with the query." + err.Error())
	}

	var dogs [][]string
	for rows.Next() {
		var id int
		var name string
		var race string
		var age int

		err := rows.Scan(&id, &name, &race, &age)
		if err != nil {
			log.Fatal(err)
		}

		dogs = append(dogs, []string{strconv.Itoa(id), name, race, strconv.Itoa(age)})
	}

	return dogs
}

func AddDogs(name string, race string, age int) {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("INSERT INTO dogs (name, race, age) VALUES (?, ?, ?)")
	statement.Exec(name, race, age)
}

func UpdateDogs(id int, name string, date string, age int) {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("UPDATE dogs SET name = ?, race = ?, age = ? WHERE id = ?;")
	statement.Exec(name, date, age, id)
}

func DeleteDogs(id int) {
	db, err := sql.Open(DBTYPE, DBPATH)
	if err != nil {
		log.Fatal("There was an error connecting to the database.")
	}
	defer db.Close()

	statement, _ := db.Prepare("DELETE FROM activities WHERE id = ?;")
	statement.Exec(id)
}
