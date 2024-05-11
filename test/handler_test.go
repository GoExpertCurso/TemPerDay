package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/GoExpertCurso/TemPerDay/internal/infra/web"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func TestSearchZipCodeValid(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/zipcode/70070550", nil)
	req = mux.SetURLVars(req, map[string]string{"cep": "70070550"})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.SearchZipCode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"temp_c":22,"temp_f":71.6,"temp_k":295.15}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSearchZipNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/zipcode/70070550", nil)
	req = mux.SetURLVars(req, map[string]string{"cep": "70070550"})
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.SearchZipCode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Location not found`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSearchZipCodeInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/zipcode/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(web.SearchZipCode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnprocessableEntity)
	}

	expected := `invalid zipcode`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
