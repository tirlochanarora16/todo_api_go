package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world 123"))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello world")
}
