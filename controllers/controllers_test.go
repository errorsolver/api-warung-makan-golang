package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/joho/godotenv"

	"golang-wm-api/controllers"
	"golang-wm-api/models"
)

// type CreateUserData struct {
// 	username, password, result string
// }
// type CreateProductData struct {
// 	productName, description, result string
// 	price                            uint32
// }

// func testCreateUser(t *testing.T) {
// 	testUserData := []CreateUserData{
// 		{"User1", "User1", "message: Success create user"},
// 		{"User2", "User2", "message: Success create user"},
// 		{"User3", "User3", "message: Success create user"},
// 	}

// 	for _, datum := range testUserData {
// 		res := CreateUser(datum.username, datum.password)

// 		if res != datum.result {
// 			t.Error("Create(%s, %s) FAIL. Expected %s got %s\n",
// 				datum.username, datum.password, datum.result, res)
// 		}
// 	}
// }

// func TestGetProducts(t *testing.T) {
// 	server := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	w := httptest.NewRecorder()

// 	GetProducts()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("expected 200 nut got %d", res.StatusCode)
// 	}
// }

func TestGetProducts(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("env error:", err)
	}
	models.ConnectDatabase()
	CC := controllers.ControllerCollection{}
	// Setup
	router := gin.Default()
	router.GET("/products", CC.GetProducts)

	// Membuat request ke endpoint /user/1
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifikasi hasil
	assert.Equal(t, http.StatusOK, w.Code)
	// Verifikasi lebih lanjut jika diperlukan
	var response struct {
		Message string           `json:"message"`
		Data    []models.Product `json:"data"`
	}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
	}
}
func TestCreateProduct(t *testing.T) {
	jsonData := []byte(`{
		"product_name":"Barang 10",
		"description":"Ini deskripsi barang 10",
		"price":1000
	}`)

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("env error:", err)
	}
	models.ConnectDatabase()
	CC := controllers.ControllerCollection{}
	router := gin.Default()

	router.POST("/product", CC.GetProducts)

	// testProductData := []CreateProductData{
	// 	{"User1", "User1", "message: Success create user", 1000},
	// }

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("Fail to make request to http %v", req)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	req.Header.Set("Content-Type", "application/json")

	var response struct {
		Message string         `json:"message"`
		Data    models.Product `json:"data"`
	}
	if err = json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Errorf("Fail to parse response JSON\nError: %v", err)
	}

	if response.Data.ID == 0 {
		t.Errorf("Expexted non-zero ID, probably data not written")
	}
	if response.Data.ProductName != "Barang 10" {
		t.Errorf("Expected product name: Barang 1, but got: %v", response.Data.ProductName)
	}
	if response.Data.Price != 100 {
		t.Errorf("Expected price 100, but got %d", response.Data.Price)
	}
}
