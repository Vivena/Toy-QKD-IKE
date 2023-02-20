package crypto

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type QKD struct {
	url        string
	port       string
	saeID      string
	HTTPClient *http.Client
}

//TODO: add mTLS
func NewQKD(url string, port string, saeID string) *QKD {

	return &QKD{url: url,
		port:  port,
		saeID: saeID,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (qkd *QKD) GetIP() string {
	return qkd.url
}

func (qkd *QKD) GetPort() string {
	return qkd.port
}

func (qkd *QKD) GetSAE_ID() string {
	return qkd.saeID
}

func (qkd *QKD) sendRequest(req *http.Request) (*Keys, error) {
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := qkd.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	var keys *RequestObj

	err = json.NewDecoder(res.Body).Decode(&keys)
	if err != nil {
		return nil, err
	}

	return &keys.Keys[0], nil

}

func (qkd *QKD) GetKey(ctx context.Context, size int) (*Keys, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s:%s/api/v1/keys/%s/enc_keys?number=1&size=%d",
		qkd.url, qkd.port, qkd.saeID, size), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	if keys, err := qkd.sendRequest(req); err != nil {
		return nil, err
	} else {
		return keys, nil
	}

}

func (qkd *QKD) GetKeyWithID() {

}
