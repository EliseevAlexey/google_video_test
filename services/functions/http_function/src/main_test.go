package main

import (
	"net/http/httptest"
	"testing"
)

func TestHttpFunction(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	HttpFunction(rr, req)

	expected := "Hello World: Http function\n"
	if actual := rr.Body.String(); actual != expected {
		t.Errorf("TestHttpFunction2. expected: %q, actual: %q", actual, expected)
	}
}
