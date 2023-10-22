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

func SendResponseData(w http.ResponseWriter, data interface{}) {

	res := ResponseData{
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
