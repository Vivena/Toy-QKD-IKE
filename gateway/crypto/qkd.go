package crypto

import (
	"context"
	"encoding/hex"
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
	SaeID      uint16
	HTTPClient *http.Client
}

type Keys struct {
	Key_id  string `json:"key_id"`
	Key_tmp string `json:"key"`
	Key     []byte
}

type RequestObj struct {
	Keys []Keys `json:"Keys"`
}

//TODO: add mTLS
func NewQKD(url string, port string, saeID uint16) *QKD {

	return &QKD{url: url,
		port:  port,
		SaeID: saeID,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// func (qkd *QKD) IP() string {
// 	return qkd.url
// }

// func (qkd *QKD) Port() string {
// 	return qkd.port
// }

// func (qkd *QKD) SAE_ID() string {
// 	return qkd.saeID
// }

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

		return nil, errors.New("unknown error")
	}

	var keys *RequestObj

	err = json.NewDecoder(res.Body).Decode(&keys)
	if err != nil {
		return nil, err
	}

	result := &keys.Keys[0]
	result.Key, err = hex.DecodeString(result.Key_tmp)
	if err != nil {
		return nil, err
	}
	// TODO: use unsafe and cGO to convert the key_tmp from hex to byte array (string is immutable in GO)

	return result, nil

}

func (qkd *QKD) GetKey(ctx context.Context, size int) (*Keys, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s/api/v1/keys/%d/enc_keys?number=1&size=%d",
		qkd.url, qkd.port, qkd.SaeID, size), nil)
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

func (qkd *QKD) GetKeyWithID(ctx context.Context, keyID string) (*Keys, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s/api/v1/keys/%s/dec_keys?key_id=%s",
		qkd.url, qkd.port, qkd.SaeID, keyID), nil)
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
