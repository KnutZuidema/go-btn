package jsonrpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Version    string        `json:"version"`
	Method     string        `json:"method"`
	Parameters []interface{} `json:"params"`
	ID         string        `json:"id"`
}

func NewRequest(method, id string, parameters ...interface{}) *Request {
	return &Request{
		Version:    "2.0",
		Method:     method,
		Parameters: parameters,
		ID:         id,
	}
}

func (r Request) ToHTTP(url string) (*http.Request, error) {
	buffer, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(buffer))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json-rpc")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept", "application/json-rpc")
	return req, nil
}
