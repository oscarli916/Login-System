package login

import (
	"fmt"
)

type otpLogin struct {
	token string
}

func (l otpLogin) Login() (userdata, error) {
	fmt.Println("Using OTP to logging in...")
	return nil, nil
}

func (l otpLogin) GetUserData() {
	fmt.Println("OTP is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l otpLogin) getUserToken() string {
	return l.token
}
