package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go-crud/models"
)

type AuthHandler struct {
	// In a real application, you would inject a user service here
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// In a real application, you would validate against database
	// This is just for demonstration
	if loginReq.Username != "admin" || loginReq.Password != "password" {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = loginReq.Username
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	// Sign the token
	tokenString, err := token.SignedString([]byte("your-secret-key")) // Use environment variable in production
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
