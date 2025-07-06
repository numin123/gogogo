package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gogogo/internal/model"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	return db
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func TestGinLoginHandlerGorm(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	password := "testpass"
	db.Create(&model.User{Email: "user@example.com", Password: hashPassword(password)})
	r := gin.Default()
	r.POST("/login", GinLoginHandlerGorm(db))

	t.Run("success", func(t *testing.T) {
		body, _ := json.Marshal(loginRequest{Email: "user@example.com", Password: password})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 200 {
			t.Errorf("[FAIL] expected 200, got %d", w.Code)
		}
		var resp map[string]string
		json.Unmarshal(w.Body.Bytes(), &resp)
		if resp["token"] == "" {
			t.Error("[FAIL] expected token in response")
		} else {
			t.Logf("[PASS] success: got token: %s", resp["token"])
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		body, _ := json.Marshal(loginRequest{Email: "user@example.com", Password: "wrong"})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 401 {
			t.Errorf("[FAIL] expected 401, got %d", w.Code)
		} else {
			t.Log("[PASS] invalid password: got 401 as expected")
		}
	})

	t.Run("invalid email", func(t *testing.T) {
		body, _ := json.Marshal(loginRequest{Email: "notfound@example.com", Password: password})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 401 {
			t.Errorf("[FAIL] expected 401, got %d", w.Code)
		} else {
			t.Log("[PASS] invalid email: got 401 as expected")
		}
	})

	t.Run("bad request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != 400 {
			t.Errorf("[FAIL] expected 400, got %d", w.Code)
		} else {
			t.Log("[PASS] bad request: got 400 as expected")
		}
	})
}
