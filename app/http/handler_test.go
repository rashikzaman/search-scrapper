package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func TestGettingUnAuthorizedUser(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestGettingIllAuthorizedUser(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/", nil)
	req.Header.Set("Authorization", "Bearer 1000")
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestRegisteringUserWithWrongInput(t *testing.T) {
	router := SetupRouter()
	pr := &UserAuthForm{
		Email:    "",
		Password: "",
	}
	body, _ := json.Marshal(pr)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/signup", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestRegisteringUserWithCorrectInput(t *testing.T) {
	router := SetupRouter()
	pr := &UserAuthForm{
		Email:    "test@gmail.com", //this will fail second time cause of email uniqueness, it's better to use seperate database for e2e test
		Password: "12345",
	}
	body, _ := json.Marshal(pr)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/signup", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestLogin(t *testing.T) {
	router := SetupRouter()
	pr := &UserAuthForm{
		Email:    "test@gmail.com", //this will fail second time cause of email uniqueness, it's better to use seperate database for e2e test
		Password: "12345",
	}
	body, _ := json.Marshal(pr)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
