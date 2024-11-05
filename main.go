package main

import (
	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/todo_api_go/database"
	"log"
	"net/http"
)

func main() {
	// load env variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables!")
		return
	}

	// establishing DB connection
	err = database.EstablishDBConnection()

	// checking for err while connecting to DB
	if err != nil {
		log.Fatal("error establishing DB connection")
		return
	}

	log.Println("DB connected successfully")

	// run DB migrations
	err = database.RunAutomigration()

	if err != nil {
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend Worked!"))
	})

	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
