package aligo

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/debuconnor/aligo-sender/aligo/object"
)

func SendMessage(data *object.AligoMessage) object.AligoResponse {
	const URL = "https://apis.aligo.in/send/"

	payload := url.Values{}
	payload.Set("key", data.Key)
	payload.Set("user_id", data.User_id)
	payload.Set("sender", data.Sender)
	payload.Set("receiver", data.Receiver)
	payload.Set("msg", data.Msg)
	payload.Set("msg_type", data.Msg_type)
	payload.Set("destination", data.Destination)
	payload.Set("rdate", data.Rdate)
	payload.Set("rtime", data.Rtime)
	payload.Set("testmode_yn", data.Testmode_yn)

	c := &http.Client{}
	req, err := c.PostForm(URL, payload)
	checkError(err)

	defer req.Body.Close()

	respBody, err := io.ReadAll(req.Body)
	checkError(err)

	var resp map[string]interface{}
	result := object.AligoResponse{}

	if err := json.Unmarshal(respBody, &resp); err == nil {
		codeType := reflect.TypeOf(resp["result_code"]).String()
		if codeType == "string" {
			result.Result_code = resp["result_code"].(string)
		} else if codeType == "float64" {
			result.Result_code = strconv.FormatFloat(resp["result_code"].(float64), 'f', -1, 64)
		}
		result.Message = resp["message"].(string)
	}

	return result
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
