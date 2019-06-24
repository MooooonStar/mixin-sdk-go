package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type P map[string]interface{}

var httpClient = &http.Client{Timeout: time.Duration(10 * time.Second)}

func BuildQuery(params P) string {
	str := make([]string, 0)
	for k, v := range params {
		str = append(str, fmt.Sprintf("%v=%v", k, v))
	}
	query := "?" + strings.Join(str, "&")
	return query
}

func RequestWithToken(method, uri string, body []byte, accessToken string) ([]byte, error) {
	req, err := http.NewRequest(method, "https://api.mixin.one"+uri, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bt, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		var resp struct {
			Error `json:"error"`
		}
		json.Unmarshal(bt, &resp)
		if resp.Code > 0 {
			err = resp.Error
		}
	}
	return bt, err
}

func Request(method, uri string, body []byte, userId, sessionId, privateKey string) ([]byte, error) {
	token, err := SignAuthenticationToken(userId, sessionId, privateKey, method, uri, string(body))
	if err != nil {
		return nil, err
	}

	return RequestWithToken(method, uri, body, token)
}
