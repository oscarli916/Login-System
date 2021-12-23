package main

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt"
	"example.com/loginSystem/login"
	"example.com/loginSystem/header"
)


var googleSigningKey = []byte("googlesigningkey")
var facebookSigningKey = []byte("facebooksigningkey")
var appleSigningKey = []byte("applesigningkey")
var OTPSigningKey = []byte("OTPsigningkey")
var errorSigningKey = []byte("errorsigningkey")



func generateToken(signingKey *[]byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString(*signingKey)
	if err != nil{
		fmt.Printf("Something went wrong: %v", err)
		return "", err
	}

	return tokenString, nil
}


func generateHeaders() ([]header.Header){
	headers := []header.Header{}

	googleToken, err := generateToken(&googleSigningKey)
	if err != nil{
		fmt.Printf("Error in generating Google Token.")
	} else{
		googleHeader := header.Header{Authorization: googleToken, Social: login.GOOGLE}
		headers = append(headers, googleHeader)
	}

	facebookToken, err := generateToken(&facebookSigningKey)
	if err != nil{
		fmt.Printf("Error in generating Facebook Token.")
	} else{
		facebookHeader := header.Header{Authorization: facebookToken, Social: login.FACEBOOK}
		headers = append(headers, facebookHeader)
	}

	appleToken, err := generateToken(&appleSigningKey)
	if err != nil{
		fmt.Printf("Error in generating Apple Token.")
	} else{
		appleHeader := header.Header{Authorization: appleToken, Social: login.APPLE}
		headers = append(headers, appleHeader)
	}

	OTPToken, err := generateToken(&OTPSigningKey)
	if err != nil{
		fmt.Printf("Error in generating OTP Token.")
	} else{
		OTPHeader := header.Header{Authorization: OTPToken, Social: login.OTP}
		headers = append(headers, OTPHeader)
	}

	errorToken, err := generateToken(&errorSigningKey)
	if err != nil{
		fmt.Printf("Error in generating Error Token.")
	} else{
		errorHeader := header.Header{Authorization: errorToken, Social: login.ERROR}
		headers = append(headers, errorHeader)
	}

	return headers

}


func main(){

	headers := generateHeaders()

	for _, header := range headers{
		login, err := login.CreateLogin(header)
		if (err != nil){
			fmt.Printf("Error in creating Login: %s \n\n", err)
		} else{
			login.Login()
			login.GetUserData()
		}
	}
}