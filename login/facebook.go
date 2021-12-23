package login

import (
	"fmt"
)

type FacebookLogin struct{
	token string
}

func (l FacebookLogin) Login(){
	fmt.Println("Using Facebook to logging in...")
}

func (l FacebookLogin) GetUserData(){
	fmt.Println("Facebook is getting user data...")
	fmt.Printf("User Token is: %v \n\n", l.getUserToken())
}

func (l FacebookLogin) getUserToken() (string){
	return l.token
}