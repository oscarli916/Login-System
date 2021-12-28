package login

type InvalidHeaderErrorHandler struct {
}

func (e InvalidHeaderErrorHandler) Error() string {
	return "Header is not supported."
}
