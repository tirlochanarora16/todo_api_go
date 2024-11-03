package main

import (
	"fmt"
	"github.com/tirlochanarora16/todo_api_go/database"
	"log"
	"net/http"
)

func main() {
	err := database.EstablishDBConnection()

	if err != nil {
		panic("error establishing DB connection")
	}

	log.Println("DB connected successfully", database.DB.Migrator().HasTable("users"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world 123"))
	})

	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world")
}
