package main

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	jwt.StandardClaims
	Email string
	Id    int
}

const (
	secret = "12345"
	issuer = "test123"
)

func genJWT(id int, email string) (string, error) {
	claims := JwtClaims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Issuer: issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Printf("%#+v", token)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(target string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(target, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}

func main() {
	// generate a jwt token
	token, err := genJWT(1, "abc123@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("generated token: ")
	fmt.Println(token)

	// decode a jwt token
	claim, err := DecodeJWT(token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%#+v\n", claim)
	fmt.Println(claim.Id)
	fmt.Println(claim.Email)
}
