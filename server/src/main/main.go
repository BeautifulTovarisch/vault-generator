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
	fmt.Println("listening..")
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:3000", routes.Routes())
}
