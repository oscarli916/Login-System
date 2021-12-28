package login

import (
	"example.com/loginSystem/header"
)

const (
	GOOGLE   = "Google"
	FACEBOOK = "Facebook"
	APPLE    = "Apple"
	OTP      = "OTP"
	ERROR    = "Error"
)

type login interface {
	Login()
	GetUserData()
}

func CreateLogin(h header.Header) (login, error) {
	if h.Social == GOOGLE {
		return googleLogin{h.Authorization}, nil
	} else if h.Social == FACEBOOK {
		return facebookLogin{h.Authorization}, nil
	} else if h.Social == APPLE {
		return appleLogin{h.Authorization}, nil
	} else if h.Social == OTP {
		return otpLogin{h.Authorization}, nil
	} else {
		return nil, InvalidHeaderErrorHandler{err: "invalid header"}
	}
}
