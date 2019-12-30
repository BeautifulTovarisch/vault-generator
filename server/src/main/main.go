package main

import (
	"fmt"
	"net/http"
	"vault-generator/routes"
)

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:2305", routes.Routes())
}
