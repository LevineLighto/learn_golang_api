package requests

import (
	"encoding/json"
	"learn_golang_api/app/helpers"
	"net/http"
)

func Read(request *http.Request, result any) {

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	helpers.PanicOnError(err)

}
