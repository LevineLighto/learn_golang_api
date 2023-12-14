package responses

import (
	"encoding/json"
	"learn_golang_api/app/helpers"
	"net/http"
)

type APIResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Write(writer http.ResponseWriter, response any) {

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	helpers.PanicOnError(err)

}
