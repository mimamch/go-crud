package serializer

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func SendResponseData(w http.ResponseWriter, status int, data interface{}) {

	res := ResponseData{
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func SendResponseMessage(w http.ResponseWriter, status int, message string) {

	res := ResponseMessage{
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
