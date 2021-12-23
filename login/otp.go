package login

import (
	"fmt"
)

type OTPLogin struct{
	token string
}

func (l OTPLogin) Login(){
	fmt.Println("Using OTP to logging in...")
}

func (l OTPLogin) GetUserData(){
	fmt.Println("OTP is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l OTPLogin) getUserToken() (string){
	return l.token
}