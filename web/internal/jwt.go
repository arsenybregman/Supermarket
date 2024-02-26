package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func UserJWT(email, password string) string {
	claims := jwt.MapClaims{
		"user_id":  123,
		"username": "johndoe",
		"exp":      time.Now().Add(time.Hour * 168).Unix(),
	}

	// Подписываем токен с помощью алгоритма HS256 и секретного ключа
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func CheckJWT() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsInVzZXJuYW1lIjoiam9obmRvZSIsImV4cCI6MTY0Njk0NzI5MX0.s487--2342342342342342342342342342342342342342342342342342342342342342"

	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return  os.Getenv("SECRET"), nil
	  })
	if err != nil {
		log.Fatal(err)
	}

	// Проверяем валидность токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Токен валиден:", claims["user_id"], claims["username"])
	} else {
		fmt.Println("Токен невалиден")
	}

}
