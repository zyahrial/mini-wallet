package controllers

import (
	"fmt"
	"log"
	"github.com/dgrijalva/jwt-go"
	"khaerus/mini-wallet/conf"
	"time"
)

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: conf.Viper("JWT_SECRET"),
		Issuer: "mini-wallet",
	}
}

type JWTService interface {
	GenerateToken(id string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct {
	secretKey string
	Issuer    string
}

type AuthCustomClaims struct {
	Id string `json:"id"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

func AuthClaim(id string) string {
	jwt_secret := conf.Viper("JWT_SECRET")
	mySigningKey := []byte(jwt_secret)

	type Auth struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := Auth{
		id,
		jwt.StandardClaims{
			ExpiresAt: 60 * 36000,
			Issuer: "mini-wallet",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Fatal(err)
	}
	return string(ss)
	// fmt.Printf("%v %v", ss, err)
}

func (service *jwtServices) GenerateToken(id string, isUser bool) string {
	claims := &AuthCustomClaims{
		id,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}