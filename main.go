package main

import (
	"fmt"
	"net/http"

	"github.com/tirlochanarora16/todo_api_go/database"
)

func main() {
	err := database.EstablishDBConnection()

	if err != nil {
		panic("error establishing DB connection")
	}

	fmt.Println("DB connected successfully")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world 123"))
	})

	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world")
}
