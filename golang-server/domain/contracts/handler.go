package contracts

import (
	"net/http"
)

type PackHandler interface {
	Calculate(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	Create(response http.ResponseWriter, request *http.Request)
	Remove(response http.ResponseWriter, request *http.Request)
}
