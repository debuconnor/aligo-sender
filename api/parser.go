package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func IsJson(contentType string) bool {
	return contentType == "application/json"
}

func DecodeJson(data *http.Request, obj interface{}) {
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(data.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&obj)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			panic("Wrong Type Field")
		} else {
			panic("Bad Request")
		}
	}
}
