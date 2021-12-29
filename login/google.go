package login

import (
	"fmt"

	jwt "github.com/golang-jwt/jwt"
)

type GoogleClaims struct {
	email          string
	email_verified string
	name           string
	picture        string
	given_name     string
	family_name    string
	locale         string
	jwt.StandardClaims
}

type googleUserData struct {
	email          string
	email_verified string
	name           string
	picture        string
	given_name     string
	family_name    string
	locale         string
}

type googleLogin struct {
	token string
}

func (l googleLogin) Login() {
	fmt.Println("Using Google to logging in...")
}

func (l googleLogin) GetUserData() {
	fmt.Println("Google is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l googleLogin) getUserToken() string {
	return l.token
}
