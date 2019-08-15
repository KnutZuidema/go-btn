package jsonrpc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
	url    string
}

func NewClient(client *http.Client, url string) *Client {
	return &Client{
		client: client,
		url:    url,
	}
}

func (c Client) Do(method string, parameters ...interface{}) (*http.Response, error) {
	httpReq, err := NewRequest(method, "id", parameters...).ToHTTP(c.url)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c Client) GetInto(dst interface{}, method string, parameters ...interface{}) (err error) {
	resp, err := c.Do(method, parameters...)
	if err != nil {
		return
	}
	defer func() {
		if err_ := resp.Body.Close(); err_ != nil {
			err = err_
		}
	}()
	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}
	var res struct {
		Id     string      `json:"id"`
		Result interface{} `json:"result"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}
	en, err := json.Marshal(res.Result)
	if err != nil {
		return
	}
	return json.Unmarshal(en, dst)
}
