package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func IsJson(contentType string) bool {
	return contentType == "application/json"
}

func DecodeJson(data *http.Request, obj interface{}) interface{} {
	var unmarshalErr *json.UnmarshalTypeError
	_obj := obj

	decoder := json.NewDecoder(data.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&_obj)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			log.Fatalln("Wrong Type Field")
		} else {
			log.Fatalln("Bad Request")
		}
	}

	return _obj
}
