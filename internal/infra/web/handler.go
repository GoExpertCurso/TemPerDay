package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GoExpertCurso/TemPerDay/internal/dto"
	"github.com/gorilla/mux"
)

func SearchZipCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cep, ok := vars["cep"]
	if !ok {
		fmt.Println("cep não encontrado")
	}

	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	/* cepResponse, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err.Error())
	} */

	var cepDto dto.Cep
	_ = json.NewDecoder(response.Body).Decode(&cepDto)
	defer response.Body.Close()
	SearchClimate(w, r, cepDto.Localidade)
}

func SearchClimate(w http.ResponseWriter, r *http.Request, location string) {
	url := "http://api.weatherapi.com/v1/current.json?key=" + "6e2adbbe3c774cf6aa122105242504" + "&q=" + location + "&aqi=yes"
	fmt.Println("URL:>", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	weatherResponse, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err.Error())
	}

	fmt.Println("Response Body:", string(weatherResponse))

	var weatherDto dto.Wheather
	_ = json.NewDecoder(response.Body).Decode(&weatherDto)
	defer response.Body.Close()
	w.Write([]byte(weatherResponse))
}
