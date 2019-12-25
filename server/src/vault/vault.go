package vault

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
)

func encrypt_config(res http.ResponseWriter, req *http.Request) {
	fmt.Println("oi")
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", encrypt_config)
	return router
}
