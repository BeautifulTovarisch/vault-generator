package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
	)

	return router
}

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
}

func main() {
	fmt.Println("listening...")
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
