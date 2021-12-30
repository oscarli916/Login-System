package login

type InvalidHeaderErrorHandler struct{}

func (e InvalidHeaderErrorHandler) Error() string {
	return "Header is not supported."
}

type EmptyTokenErrorHandler struct{}

func (e EmptyTokenErrorHandler) Error() string {
	return "Token is empty"
}

type UnsupportedIssuerErrorHandler struct{}

func (e UnsupportedIssuerErrorHandler) Error() string {
	return "Unsupported issuer"
}

type TokenExpiredErrorHandler struct{}

func (e TokenExpiredErrorHandler) Error() string {
	return "Token was expired"
}

type TokenNotIssuedErrorHandler struct{}

func (e TokenNotIssuedErrorHandler) Error() string {
	return "Token has not been issued yet"
}
