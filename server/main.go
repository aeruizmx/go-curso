package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola servidor Go!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)

	//se ejecuta el main, y en otro cmd curl localhost:8000
}
