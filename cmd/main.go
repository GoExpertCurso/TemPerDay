package main

import (
	"net/http"

	"github.com/GoExpertCurso/TemPerDay/internal/infra/web"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{cep}", web.SearchZipCode)
	http.ListenAndServe(":8080", r)
}
