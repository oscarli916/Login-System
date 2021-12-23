package login

import (
	"fmt"
)

type AppleLogin struct{
	token string
}

func (l AppleLogin) Login(){
	fmt.Println("Using Apple to logging in...")
}

func (l AppleLogin) GetUserData(){
	fmt.Println("Apple is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l AppleLogin) getUserToken() (string){
	return l.token
}