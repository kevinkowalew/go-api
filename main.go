package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	fmt.Println("Started API")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
