package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

// Creating the User Struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Creating the JWT Token Struct
type JwtToken struct {
	Token string `json:"token"`
}

// Creating the Exception Struct
type Exception struct {
	Message string `json:"message"`
}

// Creating the Response Struct
type Response struct {
	Data string `json:"data"`
}

// Creating the JWT Key
var JwtKey = []byte(os.Getenv("JWT_KEY"))

// Creating the Users Variable
var Users = []User{
	User{
		Username: "user1",
		Password: "password1",
	},
	User{
		Username: "user2",
		Password: "password2",
	},
}

// Creating the CreateToken Function
func CreateToken(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
	tokenString, error := token.SignedString(JwtKey)
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

// Creating the ProtectedEndpoint Function
func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error !")
		}
		return JwtKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode(Exception{
			Message: "Invalid authorization token",
		})
	}
}

// Creating the ValidateMiddleware Function
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return JwtKey, nil
				})

				if error != nil {
					json.NewEncoder(w).Encode(Exception{
						Message: error.Error(),
					})
					return
				}

				if token.Valid {
					next.ServeHTTP(w, r)
				} else {
					json.NewEncoder(w).Encode(Exception{
						Message: "Invalid authorization token",
					})
				}
			}
		} else {
			json.NewEncoder(w).Encode(Exception{
				Message: "An authorization header is required",
			})
		}
	})
}

// Creating the Router
/*
 * We will then add a route for our authenticate endpoint. This endpoint will be used to create JWT tokens.
 * We will then add a route for our protected endpoint. This endpoint will be used to test our JWT tokens.
 * We will then add a route for our health check endpoint. This endpoint will be used to test if our API is up and running.
 */

func main() {
	router := mux.NewRouter()
	fmt.Println("Starting the application...")
	router.HandleFunc("/authenticate", CreateToken).Methods("POST")
	router.HandleFunc("/protected", ValidateMiddleware(ProtectedEndpoint)).Methods("GET")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}

// Creating the HealthCheck Function
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Data: "API is up and running",
	})
}
