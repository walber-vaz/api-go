package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	// Cria um novo gravador de resposta HTTP falso.
	w := httptest.NewRecorder()

	// Cria uma solicitação HTTP GET falsa para o roteador.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Chama a função mainHandler para obter o roteador.
	router := mainHandler()

	// Envia a solicitação HTTP falsa para o roteador Gin.
	router.ServeHTTP(w, req)

	// Verifica o status da resposta.
	if w.Code != http.StatusOK {
		t.Errorf("status code esperado %v, mas obteve %v", http.StatusOK, w.Code)
	}

	// Verifica o corpo da resposta.
	expected := `{"message":"Hello World!"}`
	if w.Body.String() != expected {
		t.Errorf("corpo da resposta esperado %q, mas obteve %q", expected, w.Body.String())
	}
}
