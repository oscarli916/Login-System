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

type googleLogin struct {
	token string
}

func (l googleLogin) Login() (userdata, error) {
	fmt.Println("Using Google to logging in...")

	token, err := jwt.ParseWithClaims(l.getUserToken(), &GoogleClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SINGINGKEY), nil
	})
	if err != nil {
		fmt.Errorf("Cannot verify token %v\n", err)
	}

	claims, ok := token.Claims.(GoogleClaims)
	if !ok {
		fmt.Errorf("Token claims map error")
	}

	userData := map[string]interface{}{
		"email":          claims.email,
		"email_verified": claims.email_verified,
		"name":           claims.name,
		"picture":        claims.picture,
		"given_name":     claims.given_name,
		"family_name":    claims.family_name,
		"locale":         claims.locale,
	}

	return userData, nil

}

func (l googleLogin) GetUserData() {
	fmt.Println("Google is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l googleLogin) getUserToken() string {
	return l.token
}
