package main

import (
	// "fmt"

	"errors"
	"log"
	"os"

	// "net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	// "github.com/go-playground/assert"
	// "golang-wm-api/routes"
)

func TestMain(t *testing.T) {
	os.Setenv("HOST", "localhost:1234")
	godotenvLoadFunc = func() error {
		return nil
	}
	runCalled := false
	routesRunFunc = func(host string) {
		runCalled = true
	}
	connectDatabaseCalled := false
	modelsConnectDatabaseFunc = func() {
		connectDatabaseCalled = true
	}

	// Jalankan fungsi main
	main()

	// Verifikasi pemanggilan fungsi
	assert.True(t, connectDatabaseCalled)
	assert.True(t, runCalled)

	// Jalankan pengujian
	exitCode := m.Run()

	// Keluar dengan kode exit pengujian
	os.Exit(exitCode)
}

// Variabel global untuk fungsi godotenv.Load()
var godotenvLoadFunc = godotenv.Load

// Mock fungsi godotenv.Load()
func mockGodotenvLoad() error {
	return errors.New("Fail load file .env")
}

func TestMain_LoadEnvError(t *testing.T) {
	// Simulasikan lingkungan dan dependensi yang dikontrol
	os.Setenv("HOST", "localhost:8080")
	godotenvLoadFunc = mockGodotenvLoad
	loggerFatalFunc = func(v ...interface{}) {
		assert.Equal(t, "Fail load file .env", v[0])
	}

	// Jalankan fungsi main
	main()
}

// Variabel global untuk fungsi log.Fatal()
var loggerFatalFunc = log.Fatal

// Mock fungsi log.Fatal()
func mockLoggerFatal(v ...interface{}) {
	// Tidak melakukan apa pun di sini untuk menghindari aksi keluar dari pengujian
}

func TestMain_Panic(t *testing.T) {
	// Simulasikan lingkungan dan dependensi yang dikontrol
	os.Setenv("HOST", "localhost:8080")
	godotenvLoadFunc = func() error {
		panic("Panic occurred")
	}
	loggerFatalFunc = mockLoggerFatal

	// Jalankan fungsi main dan pastikan tidak ada panic
	assert.NotPanics(t, main)
}

// func TestLoadEnvFile(t *testing.T) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		t.Fatal("FAIL Get .env file", err)
// 	}
// }

// func TestMain(m *testing.M) {
// 	TestLoadEnvFile(&testing.T{})

// 	os.Exit(m.Run())
// }

// func TestGetProducts(t *testing.T) {
// 	router := routes.RoutesGroup()

// 	w := httptest.NewRecorder()
// 	req, err := http.NewRequest("GET", "/products", nil)

// 	fmt.Println("Err:", err)
// 	router.ServeHTTP(w, req)

// 	fmt.Print("error: ", w, " dan ", req)
// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, "message:Success get all products", w.Body.String())
// }

// func TestMain(t *testing.T) {
// 	url := "http://localhost:8080/api/v1/products"
// 	t.Run("TestSuccessRequest", func(t *testing.T) {
// 		httpNewReq, httpNewReqErr := http.NewRequest("GET", url, nil)
// 		if httpNewReqErr != nil {
// 			t.Error(httpNewReqErr)
// 		}

// 		// req.Header.Add("authorization", "Basic ")
// 		httpNewReq.Header.Add("cache-control", "no-cache")

// 		httpDefaultRes, httpDefaultResErr := http.DefaultClient.Do(httpNewReq)
// 		if httpDefaultResErr != nil {
// 			t.Error(httpDefaultResErr)
// 		}

// 		body, bodyErr := io.ReadAll(httpDefaultRes.Body)
// 		if bodyErr != nil {
// 			t.Error(bodyErr)
// 		}

// 		t.Log(httpDefaultRes)
// 		t.Log(string(body))
// 		defer httpDefaultRes.Body.Close()
// 	})
// }

// func TestMain(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		url            string
// 		expectedStatus int
// 	}{
// 		{
// 			name:           "TestSuccessRequest",
// 			url:            "http://localhost:8080/api/v1/products",
// 			expectedStatus: http.StatusOK,
// 		},
// 		{
// 			name:           "TestNotFoundRequest",
// 			url:            "http://localhost:8080/api/v1/nonexistent",
// 			expectedStatus: http.StatusNotFound,
// 		},
// 		{
// 			name:           "TestInvalidInputRequest",
// 			url:            "http://localhost:8080/api/v1/products?param=value",
// 			expectedStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name:           "TestUnauthorizedRequest",
// 			url:            "http://localhost:8080/api/v1/protected",
// 			expectedStatus: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			req, err := http.NewRequest("GET", tt.url, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			req.Header.Add("cache-control", "no-cache")

// 			res, err := http.DefaultClient.Do(req)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			defer res.Body.Close()

// 			if res.StatusCode != tt.expectedStatus {
// 				t.Errorf("Kode status tidak sesuai. Harapannya: %d, Dapatkan: %d", tt.expectedStatus, res.StatusCode)
// 			}

// 			body, err := io.ReadAll(res.Body)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			t.Log(string(body))
// 		})
// 	}
// }
