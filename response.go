package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	ErrorCode string      `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func (response *Response) send(w http.ResponseWriter) {
	body, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%v", string(body))
}
