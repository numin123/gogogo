package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gogogo/internal/model"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func setupUserTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	return db
}

func TestGinCreateUserHandlerGorm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupUserTestDB()
	r := gin.Default()
	r.POST("/users", GinCreateUserHandlerGorm(db))

	t.Run("success", func(t *testing.T) {
		body, _ := json.Marshal(createUserRequest{Name: "Test User", Email: "newuser@example.com", Password: "password123"})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 200 {
			t.Errorf("[FAIL] expected 200, got %d", w.Code)
		}
		var resp createUserResponse
		json.Unmarshal(w.Body.Bytes(), &resp)
		if resp.Email != "newuser@example.com" || resp.Name != "Test User" {
			t.Error("[FAIL] expected correct user in response")
		} else {
			t.Logf("[PASS] success: created user %s (%s)", resp.Name, resp.Email)
		}
	})

	t.Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 400 {
			t.Errorf("[FAIL] expected 400, got %d", w.Code)
		} else {
			t.Log("[PASS] bad request: got 400 as expected")
		}
	})

	t.Run("duplicate email", func(t *testing.T) {
		// Create user first
		db.Create(&model.User{Name: "Dup User", Email: "dup@example.com", Password: "pass"})
		body, _ := json.Marshal(createUserRequest{Name: "Dup User", Email: "dup@example.com", Password: "pass"})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 500 {
			t.Errorf("[FAIL] expected 500, got %d", w.Code)
		} else {
			t.Log("[PASS] duplicate email: got 500 as expected")
		}
	})
}
