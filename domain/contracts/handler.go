package contracts

import (
	"net/http"
)

type PackHandler interface {
	Calculate(response http.ResponseWriter, request *http.Request)
}
