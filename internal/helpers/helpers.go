package helpers 

import "net/http"

var (
    InternalServerErrorMessage = "An unexpected error happened, try again!"
    BadRequestErrorMessage = "Please insert a valid value"
)

func ServerError(w http.ResponseWriter, response []byte) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write(response)
}

func BadRequest(w http.ResponseWriter, response []byte) {
    w.WriteHeader(http.StatusBadRequest)
    w.Write(response)
}
