package product_controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Find(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
