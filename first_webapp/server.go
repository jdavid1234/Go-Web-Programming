package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Yossi David, %s!", request.URL.Path[1:1])

}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8030", nil)

}
