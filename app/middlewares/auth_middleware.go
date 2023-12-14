package middlewares

import (
	"learn_golang_api/app/responses"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if "Auwooo" == request.Header.Get("Authorization") {
		middleware.Handler.ServeHTTP(writer, request)

		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	response := responses.APIResponse{
		Message: "UNAUTHORIZED",
	}

	responses.Write(writer, response)

}
