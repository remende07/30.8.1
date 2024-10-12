package main

import (
	"fmt"
	"log"
	"sasha/Desktop/30.8.1-main/pkg/storage"
	"sasha/Desktop/30.8.1-main/pkg/storage/postgres"
)

var db storage.Interface

func main() {
	var err error

	connstr :=
		"postgres://postgres:qwerty@localhost/tasks?sslmode=disable"

	db, err = postgres.New(connstr)

	if err != nil {
		log.Fatal(err)
	}

	tasks, err := db.GetTasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tasks)
}
