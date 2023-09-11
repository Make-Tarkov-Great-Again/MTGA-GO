package webapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_APP_KEY = "training.go"

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type response struct {
	Token string `json:"token"`
}

type responseError struct {
	Error string `json:"error"`
}

func decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

func handleTokenCreate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		err := decode(w, r, &req)
		if err != nil {
			msg := fmt.Sprintf("Cannot parse login body. err=%v", err)
			log.Println(msg)
			respond(w, r, nil, http.StatusBadRequest)
			respond(w, r, responseError{
				Error: msg,
			}, http.StatusUnauthorized)
			return
		}
		// database version
		/*
			found, err := s.store.FindUser(req.Username, req.Password)
			if err != nil {
				msg := fmt.Sprintf("Cannot find user. err=%v", err)
				log.Println(msg)
				s.respond(w, r, nil, http.StatusBadRequest)
				s.respond(w, r, responseError{
					Error: msg,
				}, http.StatusInternalServerError)
				return
			}
		*/

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Email,
			"exp":      time.Now().Add(time.Hour * time.Duration(24)).Unix(), // expiration date
			"iat":      time.Now().Unix(),                                    // Issued At Time
		})

		tokenStr, err := token.SignedString([]byte(JWT_APP_KEY))
		if err != nil {
			msg := fmt.Sprintf("Cannot generate JWT. err=%v", err)
			respond(w, r, responseError{
				Error: msg,
			}, http.StatusInternalServerError)
			return
		}

		respond(w, r, response{
			Token: tokenStr,
		}, http.StatusOK)
	}
}
