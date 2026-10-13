package helper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(payload any) error {

	err := validate.Struct(payload)
	if err != nil {
		return err
	}

	return nil
}
func WriteJsonBody(
    w http.ResponseWriter,
    statusCode int,
    data interface{},
) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        return
    }
}

func WriteJsonSucResp(
    w http.ResponseWriter,
    statusCode int,
    message string,
    data interface{},
) {
    response := map[string]interface{}{
        "message": message,
        "data":    data,
    }

    WriteJsonBody(w, statusCode, response)
}

func WriteJsonErrResponse(
    w http.ResponseWriter,
    statusCode int,
    message string,
    err error,
) {
    response := map[string]interface{}{
        "message": message,
    }

    if err != nil {
        response["error"] = err.Error()
    }

    WriteJsonBody(w, statusCode, response)
}
func ReadJsonBody(r *http.Request, payload any) error {

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		return err
	}

	return nil
}