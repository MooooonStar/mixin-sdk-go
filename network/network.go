package network

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"time"
)

func CreatePIN(old_pin, new_pin, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	var oldEncryptedPin string
	if len(old_pin) > 0 {
		oldEncryptedPin = EncryptPIN(old_pin, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	}
	newEncryptedPin := EncryptPIN(new_pin, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	params := P{"old_pin": oldEncryptedPin, "pin": newEncryptedPin}
	return MixinRequest("POST", "/pin/update", params, userId, sessionId, privateKey)
}

func VerifyPIN(pin, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	encryptedPin := EncryptPIN(pin, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	body := P{"pin": encryptedPin}
	return MixinRequest("POST", "/pin/verify", body, userId, sessionId, privateKey)
}

func Deposit(assetID string, userId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/assets/"+assetID, nil, userId, sessionId, privateKey)
}

func Withdrawal(addressId, amount, memo, trace string, pinCode, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"address_id": addressId,
		"amount":     amount,
		"pin":        EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano())),
		"trace_id":   trace,
		"memo":       memo,
	}
	return MixinRequest("POST", "/withdrawals", params, userId, sessionId, privateKey)
}

func CreateAddress(assetID, publicOrName, emptyOrTag string, pinCode, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	pin := EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	params := P{"asset_id": assetID, "pin": pin}
	if emptyOrTag != "" {
		params["account_name"] = publicOrName
		params["account_tag"] = emptyOrTag
	} else {
		params["public_key"] = publicOrName
		params["label"] = emptyOrTag
	}
	return MixinRequest("POST", "/addresses", params, userId, sessionId, privateKey)
}

func DeleteAddress(addressID string, pinCode, pinToken, usedId, sessionId, privateKey string) ([]byte, error) {
	uri := "/addresses/" + addressID + "/delete"
	pin := EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	params := P{"pin": pin}
	return MixinRequest("POST", uri, params, usedId, sessionId, privateKey)
}

func ReadAddress(addressId string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/addresses/"+addressId, nil, usedId, sessionId, privateKey)
}

func WithdrawalAddresses(assetID string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/assets/"+assetID+"/addresses", nil, usedId, sessionId, privateKey)
}

func ReadAsset(assetID string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/assets/"+assetID, nil, usedId, sessionId, privateKey)
}

func ReadAssets(usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/assets", nil, usedId, sessionId, privateKey)
}

func VerifyPayment(opponentId, amount, assetId, traceId string, usedId, sessionId, privateKey string) ([]byte, error) {
	params := P{
		"asset_id":    assetId,
		"opponent_id": opponentId,
		"amount":      amount,
		"trace_id":    traceId,
	}
	return MixinRequest("POST", "/payments", params, usedId, sessionId, privateKey)
}

func Transfer(opponentID, amount, asset, memo, trace string, pinCode, pinToken, userId, sessionId, privateKey string) ([]byte, error) {
	pin := EncryptPIN(pinCode, pinToken, sessionId, privateKey, uint64(time.Now().UnixNano()))
	params := P{
		"asset_id":    asset,
		"opponent_id": opponentID,
		"amount":      amount,
		"pin":         pin,
		"trace_id":    trace,
		"memo":        memo,
	}
	return MixinRequest("POST", "/transfers", params, userId, sessionId, privateKey)
}

func ReadTransfer(traceId string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/transfers/trace/"+traceId, nil, usedId, sessionId, privateKey)
}

func TopAssets(usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/network/assets/top", nil, usedId, sessionId, privateKey)
}

func NetworkAsset(assetID, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/network/assets/"+assetID, nil, usedId, sessionId, privateKey)
}

func SearchAssets(symbol string, usedId, sessionId, privateKey string) ([]byte, error) {
	return MixinRequest("GET", "/network/assets/search/"+symbol, nil, usedId, sessionId, privateKey)
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
