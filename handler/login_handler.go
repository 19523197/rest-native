package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"rest-native/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"Username"`
}

type UserHandler struct {
	Repo repository.UserRepository
}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Invalid credential", http.StatusBadRequest)
		return
	}

	ok, userInfo := u.Repo.AuthenticateUser(username, password)
	if !ok {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	claims := JWTClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APPLICATION_NAME"),
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
		Username: userInfo["username"].(string),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString("sadwawduu211ud8hd1i")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenString, _ := json.Marshal(signedToken)
	w.Write(tokenString)

}
