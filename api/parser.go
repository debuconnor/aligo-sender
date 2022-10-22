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

func DecodeJson(data *http.Request, obj struct{}) (result int) {
	var unmarshalErr *json.UnmarshalTypeError
	_obj := obj
	result = -1

	log.Println(obj)
	decoder := json.NewDecoder(data.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&_obj)
	log.Println(_obj)
	if err != nil {
		/*
			Err list
			1 : Wrong Type Field
			2 : Bad Request
		*/
		if errors.As(err, &unmarshalErr) {
			result = 1
		} else {
			result = 2
		}
	}

	return
}
