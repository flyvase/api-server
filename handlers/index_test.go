package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/?name=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"Unexpected status: got (%v) want (%v)",
			status,
			http.StatusOK,
		)
	}

	expected := "Hello test"
	if rr.Body.String() != expected {
		t.Errorf(
			"Unexpected body: get (%v) want (%v)",
			rr.Body.String(),
			expected,
		)
	}
}
