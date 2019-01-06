package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/json-iterator/go"
)

type P map[string]interface{}

const (
	baseUrl = "https://api.mixin.one"
)

var client *http.Client

func init() {
	client = &http.Client{Timeout: time.Duration(30 * time.Second)}
}

func Request(method, uri string, body []byte, clientId, sessionId, privateKey string) ([]byte, error) {
	token, err := SignAuthenticationToken(clientId, sessionId, privateKey, method, uri, string(body))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseUrl+uri, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bt, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var resp struct {
			Error Error `json:"error"`
		}
		err = json.Unmarshal(bt, &resp)
		if err == nil {
			err = resp.Error
		}
	}
	return bt, err
}

func MixinRequest(method, uri string, params ...P) ([]byte, error) {
	if len(params) == 0 {
		return Request(method, uri, nil, UserId, SessionId, PrivateKey)
	}

	switch method {
	case "GET":
		str := make([]string, 0)
		for k, v := range params[0] {
			str = append(str, fmt.Sprintf("%v=%v", k, v))
		}
		query := "?" + strings.Join(str, "&")
		return Request(method, uri+query, nil, UserId, SessionId, PrivateKey)

	case "POST":
		body, err := jsoniter.Marshal(params[0])
		if err != nil {
			return nil, err
		}
		return Request(method, uri, body, UserId, SessionId, PrivateKey)
	}
	return nil, fmt.Errorf("Unsupported method.")
}