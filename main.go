package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := ""

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
