package network

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func NewUser(userId, sessionId, pinToken, privateKey string) *User {
	return &User{UserId: userId, SessionId: SessionId, PinToken: pinToken, PrivateKey: privateKey}
}

func (u User) CreatePIN(old_pin, new_pin string) ([]byte, error) {
	method := "POST"
	uri := "/pin/update"

	oldEncryptedPin := ""
	if len(old_pin) > 0 {
		oldEncryptedPin = EncryptPIN(old_pin, u.PinToken, u.SessionId, u.PrivateKey, uint64(time.Now().UnixNano()))
	}

	newEncryptedPin := EncryptPIN(new_pin, u.PinToken, u.SessionId, u.PrivateKey, uint64(time.Now().UnixNano()))

	body := P{"old_pin": oldEncryptedPin, "pin": newEncryptedPin}
	return u.MixinRequest(method, uri, body)
}

func (u User) VerifyPIN(pin string) ([]byte, error) {
	method := "POST"
	uri := "/pin/verify"
	encryptedPin := EncryptPIN(pin, u.PinToken, u.SessionId, u.PrivateKey, uint64(time.Now().UnixNano()))

	body := P{"pin": encryptedPin}
	return u.MixinRequest(method, uri, body)
}

func (u User) Deposit(asset string, acount_info ...string) ([]byte, error) {
	method := "GET"
	uri := "/assets/" + asset
	if asset == EOS {
		if len(acount_info) == 2 {
			body := P{"account_name": acount_info[0], "account_tag": acount_info[1]}
			return u.MixinRequest(method, uri, body)
		}
	}

	return u.MixinRequest(method, uri)
}

func (u User) Withdrawal(asset, address, amount string, name ...string) ([]byte, error) {
	method := "POST"
	uri := "/withdrawals"

	label := "Mixin " + asset + " Address"
	if len(name) > 0 {
		label = name[0]
	}
	data, err := u.CreateAddress(asset, address, label)
	if err != nil {
		return nil, err
	}

	type Resp struct {
		Data struct {
			AddressId string `json:"address_id"`
		} `json:"data"`
	}

	var resp Resp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	addressId := resp.Data.AddressId
	body := P{
		"address_id": addressId,
		"amount":     amount,
		"pin":        EncryptPIN(u.PinCode, u.PinToken, u.SessionId, u.PrivateKey, uint64(time.Now().UnixNano())),
		"trace_id":   uuid.Must(uuid.NewV4()).String(),
		"memo":       "Created By Mibot",
	}

	return u.MixinRequest(method, uri, body)
}

func (u User) CreateAddress(asset, address, label string, account_info ...string) ([]byte, error) {
	method := "POST"
	uri := "/addresses"

	pin := EncryptPIN(u.PinCode, u.PinToken, u.SessionId, u.PrivateKey, uint64(time.Now().UnixNano()))

	body := P{
		"asset_id":   asset,
		"public_key": address,
		"label":      label,
		"pin":        pin,
	}

	if asset == EOS {
		if len(account_info) == 2 {
			body["account_name"] = account_info[0]
			body["account_tag"] = account_info[1]
		}
	}
	return u.MixinRequest(method, uri, body)
}

func (u User) ReadAsset(asset string) ([]byte, error) {
	method := "GET"
	uri := "/assets/" + asset
	return u.MixinRequest(method, uri)
}

func (u User) ReadAssets() ([]byte, error) {
	method := "GET"
	uri := "/assets"
	return u.MixinRequest(method, uri)
}

func (u User) VarifyPayment(opponent_id, amount, asset, trace_id string) ([]byte, error) {
	method := "POST"
	uri := "/payments"
	body := P{
		"asset_id":    asset,
		"opponent_id": opponent_id,
		"amount":      amount,
		"trace_id":    trace_id,
	}
	return u.MixinRequest(method, uri, body)
}

func (u User) Transfer(opponent_id, amount, asset, memo string, trace_id ...string) ([]byte, error) {
	method := "POST"
	uri := "/transfers"
	pin := EncryptPIN(u.PinCode, u.PinToken, u.SessionId, u.PrivateKey, uint64(time.Now().UnixNano()))

	trace_uuid := uuid.Must(uuid.NewV4()).String()
	if len(trace_id) > 0 {
		trace_uuid = trace_id[0]
	}
	body := P{
		"asset_id":    asset,
		"opponent_id": opponent_id,
		"amount":      amount,
		"pin":         pin,
		"trace_id":    trace_uuid,
		"memo":        memo,
	}
	return u.MixinRequest(method, uri, body)
}

func (u User) ReadProfile() ([]byte, error) {
	method := "GET"
	uri := "/me"
	return u.MixinRequest(method, uri)
}

func (u User) Request(method, uri string, body []byte) ([]byte, error) {
	return Request(method, uri, body, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) MixinRequest(method, uri string, payload ...P) ([]byte, error) {
	if len(payload) == 0 {
		return u.Request(method, uri, nil)
	}

	switch method {
	case "GET":
		str := make([]string, 0)
		for k, v := range payload[0] {
			str = append(str, fmt.Sprintf("%v=%v", k, v))
		}
		query := "?" + strings.Join(str, "&")
		return u.Request(method, uri+query, nil)

	case "POST":
		bt, err := json.Marshal(payload[0])
		if err != nil {
			return nil, err
		}
		return u.Request(method, uri, bt)
	}
	return nil, fmt.Errorf("Unsupported method.")
}
