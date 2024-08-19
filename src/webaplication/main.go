package main

import (
	"log"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	message := []byte(`
	<h1>Hello, web!</h1>
	<hr />
	<h3>How are you doing</h3>
	`)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
