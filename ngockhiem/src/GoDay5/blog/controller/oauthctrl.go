package controller

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type Token string

const (
	TOKEN_LENGTH = 99999
)

var oauthToken = make(map[*http.Request]Token)

func generateToken() Token {
	token := Token("Token" + strconv.Itoa(rand.Intn(TOKEN_LENGTH)))
	return token
}

func GetToken(r *http.Request) (Token, error) {
	if token, ok := oauthToken[r]; ok {
		return token, nil
	}
	err := errors.New("Token not found")
	return Token(""), err
}

// /login
func OauthGet(w http.ResponseWriter, r *http.Request) {
	if token, ok := oauthToken[r]; ok {
		w.WriteHeader(200)
		fmt.Fprintf(w, string(token))
		return
	}
	token := generateToken()
	w.WriteHeader(200)
	fmt.Fprintf(w, string(token))
}

// /login/reset
func OauthReset(w http.ResponseWriter, r *http.Request) {
	token := Token("Token" + strconv.Itoa(rand.Intn(99999)))
	fmt.Fprintf(w, string(token))
}
