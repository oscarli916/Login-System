package login

import (
	"fmt"
)

type GoogleLogin struct{
	token string
}

func (l GoogleLogin) Login(){
	fmt.Println("Using Google to logging in...")
}

func (l GoogleLogin) GetUserData(){
	fmt.Println("Google is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l GoogleLogin) getUserToken() (string){
	return l.token
}