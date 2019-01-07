package network

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CreatePIN(old_pin, new_pin, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	method := "POST"
	uri := "/pin/update"

	oldEncryptedPin := EncryptPIN(old_pin, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	newEncryptedPin := EncryptPIN(new_pin, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))

	body := P{
		"old_pin": oldEncryptedPin,
		"pin":     newEncryptedPin,
	}

	return MixinRequest(method, uri, body, userId, sessionId, privateKey)
}

func VerifyPIN(pin, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	method := "POST"
	uri := "/pin/verify"
	encryptedPin := EncryptPIN(pin, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))

	body := P{
		"pin": encryptedPin,
	}

	return MixinRequest(method, uri, body, userId, sessionId, privateKey)
}

func Deposit(symbol string, userId, sessionId, privateKey string, acountInfo ...string) ([]byte, error) {
	method := "GET"
	uri := "/assets/" + symbolAssetId[symbol]
	if symbol == EOS {
		if len(acountInfo) == 2 {
			body := P{
				"account_name": acountInfo[0],
				"account_tag":  acountInfo[1],
			}
			return MixinRequest(method, uri, body, userId, sessionId, privateKey)
		}
	}

	return MixinRequest(method, uri, nil, userId, sessionId, privateKey)
}

func Withdrawal(addressId, amount, memo string, pinCode, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	method := "POST"
	uri := "/withdrawals"

	body := P{
		"address_id": addressId,
		"amount":     amount,
		"pin":        EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano())),
		"trace_id":   uuid.Must(uuid.NewV4()).String(),
		"memo":       memo,
	}

	return MixinRequest(method, uri, body, userId, sessionId, privateKey)
}

func CreateAddress(symbol, address, label string, pinCode, pinToken, userId, sessionId, privateKey string, acountInfo ...string) ([]byte, error) {
	method := "POST"
	uri := "/addresses"

	pin := EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))

	body := P{
		"asset_id":   symbolAssetId[symbol],
		"public_key": address,
		"label":      label,
		"pin":        pin,
	}

	if symbol == EOS {
		if len(acountInfo) == 2 {
			body["account_name"] = acountInfo[0]
			body["account_tag"] = acountInfo[1]
		}
	}
	return MixinRequest(method, uri, body, userId, sessionId, privateKey)
}

func DeleteAddress(id string, pinCode, pinToken, usedId, sessionId, privateKey string) ([]byte, error) {
	method := "POST"
	uri := "/addresses/" + id + "/delete"
	pin := EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))

	body := P{
		"pin": pin,
	}
	return MixinRequest(method, uri, body, usedId, sessionId, privateKey)
}

func WithdrawalAddresses(symbol string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/assets/"+symbolAssetId[symbol]+"/addresses", nil, usedId, sessionId, privateKey)
}

func ReadAddress(addressId string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/addresses/"+addressId, nil, usedId, sessionId, privateKey)
}

func ReadAsset(symbol string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/assets/"+symbolAssetId[symbol], nil, usedId, sessionId, privateKey)
}

func ReadAssets(usedId, sessionId, privateKey string) ([]byte, error) {
	method := "GET"
	uri := "/assets"
	return MixinRequest(method, uri, nil, usedId, sessionId, privateKey)
}

func VarifyPayment(opponent_id, amount, symbol, traceId string, usedId, sessionId, privateKey string) ([]byte, error) {
	method := "POST"
	uri := "/payments"
	body := P{
		"asset_id":    symbolAssetId[symbol],
		"opponent_id": opponent_id,
		"amount":      amount,
		"trace_id":    traceId,
	}
	return MixinRequest(method, uri, body, usedId, sessionId, privateKey)
}

func Transfer(opponent_id, amount, symbol, memo string, pinCode, pinToken, userId, sessionId, privateKey string, traceId ...string) ([]byte, error) {
	method := "POST"
	uri := "/transfers"
	pin := EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))

	trace_uuid := uuid.Must(uuid.NewV4()).String()
	if len(traceId) > 0 {
		trace_uuid = traceId[0]
	}
	body := P{
		"asset_id":    symbolAssetId[symbol],
		"opponent_id": opponent_id,
		"amount":      amount,
		"pin":         pin,
		"trace_id":    trace_uuid,
		"memo":        memo,
	}

	return MixinRequest(method, uri, body, userId, sessionId, privateKey)
}

func ReadTransfer(traceId string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/transfers/trace/"+traceId, nil, usedId, sessionId, privateKey)
}

func TopAssets(usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/network", nil, usedId, sessionId, privateKey)
}

func SearchAssets(query string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/network/assets/search/"+query, nil, usedId, sessionId, privateKey)
}

func CreateAppUser(name, pin string, usedId, sessionId, privateKey string) (*User, error) {
	private, _ := rsa.GenerateKey(rand.Reader, 1024)
	block := x509.MarshalPKCS1PrivateKey(private)
	b := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: block,
	}
	bt := pem.EncodeToMemory(b)

	publicKey, err := x509.MarshalPKIXPublicKey(private.Public())
	if err != nil {
		return nil, err
	}

	body := P{
		"session_secret": base64.StdEncoding.EncodeToString(publicKey),
		"full_name":      name,
	}
	data, err := MixinRequest("POST", "/users", body, usedId, sessionId, privateKey)
	if err != nil {
		return nil, err
	}

	var Resp struct {
		Data User `json:"data"`
	}
	err = json.Unmarshal(data, &Resp)
	if err != nil {
		return nil, err
	}

	user := Resp.Data
	user.PrivateKey = string(bt)

	_, err = user.CreatePIN("", pin)
	if err != nil {
		return nil, err
	}
	user.PinCode = pin
	return &user, err
}
