package main

import (
	"net/http"

	"github.com/GoExpertCurso/TemPerDay/configs"
	"github.com/GoExpertCurso/TemPerDay/internal/infra/web"
	"github.com/gorilla/mux"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/{cep}", web.SearchZipCode)
	http.ListenAndServe(":"+configs.WEB_SERVER_PORT, r)
}
