package login

import (
	"fmt"
)

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
