package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{cep}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cep, ok := vars["cep"]
		if !ok {
			fmt.Println("cep não encontrado")
		}
		fmt.Fprintf(w, "cep: %s", cep)
	})
	http.ListenAndServe(":8080", r)
}
