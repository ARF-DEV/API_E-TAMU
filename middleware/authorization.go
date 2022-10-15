package middleware

import (
	"E-TamuAPI/helpers"
	"E-TamuAPI/repository"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Authorization(userRepo *repository.UserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("Authorization")

			if !strings.Contains(bearer, "Bearer") {
				log.Println("Error : Bearer Token Not Found")
				helpers.ErrorResponseJSON(w, "Bearer Token Not Found", http.StatusBadRequest)
				return
			}

			tokenString := strings.Replace(bearer, "Bearer ", "", -1)
			tokenHeader := strings.Split(tokenString, ".")[0]
			decotedByte, err := base64.StdEncoding.DecodeString(tokenHeader)
			if err != nil {
				log.Println("Failed to decode base64: ", err.Error())
				helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
				return
			}
			var tokenHeaderMap map[string]interface{}

			json.Unmarshal(decotedByte, &tokenHeaderMap)
			if tokenHeaderMap["typ"].(string) != "JWT" {
				log.Println("Token type doesn't match")
				helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
				return
			}

			userClaims := helpers.UserClaims{}
			token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("signing method invalid")
				} else if method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("signing method invalid")
				}

				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})
			if err != nil {
				log.Println("Error on parsing claim: ", err.Error())
				helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				log.Println("Invalid Token")
				helpers.ErrorResponseJSON(w, "Invalid Token", http.StatusUnauthorized)
				return
			}
			user := userClaims.UserData

			ctx := context.WithValue(r.Context(), "user_data", user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
