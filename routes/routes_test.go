package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPingRoute(t *testing.T) {
// 	router := RoutesCollection()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/ping", nil)
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "pong", w.Body.String())
// }

func TestGetProducts(t *testing.T) {
	router := RoutesGroup()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/products", nil)

	fmt.Println("Err:", err)
	router.ServeHTTP(w, req)

	fmt.Print("error: ", w, " dan ", req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}
