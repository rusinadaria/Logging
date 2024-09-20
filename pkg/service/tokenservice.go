package main

import (
	_"github.com/lib/pq"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rusinadaria/Logging/repository"
)

// type tokenService interface {
// 	generateTokens(guid int, ip string) (string, string)
// 	saveTokens()
// }

// type tokens struct {
// 	accessToken string
// 	refreshToken string
// }

func generateTokens(ip string) (string, string) {
	accessToken, err := generateAccessToken(ip)
	if err != nil {
		fmt.Println("tokenService error")
	}
	fmt.Printf("User Access token: %s\n", accessToken)

	refreshToken, err := generateRefreshToken(ip)
	if err != nil {
		fmt.Println("tokenService error")
	}
	fmt.Printf("User Refresh token: %s\n", refreshToken)

	return accessToken, refreshToken
}

func generateAccessToken (ip string) (string, error) {
	secret := []byte("secret_key") // брать из env
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_ip": ip,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString(secret)
}

func generateRefreshToken (ip string) (string, error) {
	secret := []byte("secret_key") // брать из env
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_ip": ip,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secret)
}

func saveToken(refreshToken string, guid int) {
	SaveRefreshToken(guid, refreshToken)
}

func validTokens () { // для второго маршрута
	//распарсить токен

	//сравнить refresh токен из бд с полученным от клиента
	//проверить, не изменился ли ip из payload
	// если изменился, отправить письмо на email, используя моковые данные


}

// func sendActivationMail() {
	
// }

// func refreshTokens () { // для второго маршрута
	// после валидации сгенерировать новую пару токенов
// }