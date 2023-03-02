package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"task-5-vix-btpns-SuburFirmansyah/models"
	"time"
)

func GenerateJwt(user models.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetEnv("JWT_SECRET")))
}

func TokenValidation(r *http.Request, db *gorm.DB) int {
	tokenString := getTokenHeader(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		return http.StatusUnauthorized
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"]
		if userId == nil {
			return http.StatusUnauthorized
		}

		var user models.User
		_, err := user.GetUser(db, uint(userId.(float64)))
		if err != nil {
			return http.StatusUnauthorized
		}
	}
	return http.StatusOK
}

func getTokenHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	bearerString := "bearer "
	if len(bearerToken) < len(bearerString) {
		return ""
	}
	return bearerToken[len(bearerString):]
}

func GetUserIdFromToken(r *http.Request) (uint, error) {
	tokenString := getTokenHeader(r)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(GetEnv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["user_id"].(float64)), nil
	}
	return 0, nil
}
