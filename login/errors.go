package login

import (
	"fmt"
)

type InvalidHeaderErrorHandler struct {
	err string
}

func (e InvalidHeaderErrorHandler) Error() string {
	return fmt.Sprintf("Error type: [Invalid Header] Error message: %v", e.err)
}
