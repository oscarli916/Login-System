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

type Login interface {
	Login()
	GetUserData()
}

func CreateLogin(h header.Header) (Login, error) {
	if h.Social == GOOGLE {
		return GoogleLogin{h.Authorization}, nil
	} else if h.Social == FACEBOOK {
		return FacebookLogin{h.Authorization}, nil
	} else if h.Social == APPLE {
		return AppleLogin{h.Authorization}, nil
	} else if h.Social == OTP {
		return OTPLogin{h.Authorization}, nil
	} else {
		return nil, InvalidHeaderErrorHandler{err: "invalid header"}
	}
}
