package gee_test

import (
	"gee/gee"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEngineGet(t *testing.T) {
	engine := gee.New()
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello Get")
	}
	engine.GET("/test", testHandler)

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Experted statusOK code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestEnginePost(t *testing.T) {


}
