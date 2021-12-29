package login

import (
	"fmt"
)

type appleLogin struct {
	token string
}

func (l appleLogin) Login() (userdata, error) {
	fmt.Println("Using Apple to logging in...")
	return nil, nil
}

func (l appleLogin) GetUserData() {
	fmt.Println("Apple is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l appleLogin) getUserToken() string {
	return l.token
}
