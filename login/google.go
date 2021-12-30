package login

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

type GoogleClaims struct {
	Email          string `json:"email"`
	Email_verified string `json:"email_verified"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Locale         string `json:"locale"`
	jwt.StandardClaims
}

type googleLogin struct {
	token string
}

func (l googleLogin) Login() (userdata, error) {
	fmt.Println("Using Google to logging in...")

	token, err := jwt.ParseWithClaims(l.getUserToken(), &GoogleClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNINGKEY), nil
	})
	if err != nil {
		fmt.Errorf("Cannot verify token %v\n", err)
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		fmt.Errorf("Token claims map error")
	}

	// Hanlde empty token
	if (jwt.StandardClaims{} == claims.StandardClaims) {
		return nil, fmt.Errorf("Token is empty")
	}

	// Handle issuer
	if claims.Issuer != GOOGLEISSUER {
		return nil, fmt.Errorf("Unsupported issuer")
	}

	// Handle expired
	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		return nil, fmt.Errorf("Token was expired")
	}

	// Handle token not been issued
	if !claims.VerifyIssuedAt(time.Now().Unix(), false) {
		return nil, fmt.Errorf("Token has not been issued yet")
	}

	userData := map[string]interface{}{
		"email":          claims.Email,
		"email_verified": claims.Email_verified,
		"name":           claims.Name,
		"picture":        claims.Picture,
		"given_name":     claims.Given_name,
		"family_name":    claims.Family_name,
		"locale":         claims.Locale,
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
