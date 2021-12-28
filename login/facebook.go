package login

import (
	"fmt"
)

type facebookLogin struct {
	token string
}

func (l facebookLogin) Login() {
	fmt.Println("Using Facebook to logging in...")
}

func (l facebookLogin) GetUserData() {
	fmt.Println("Facebook is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l facebookLogin) getUserToken() string {
	return l.token
}
