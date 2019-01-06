package network

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"time"

	uuid "github.com/satori/go.uuid"
)

func CreatePIN(old_pin, new_pin string) ([]byte, error) {
	method := "POST"
	uri := "/pin/update"

	oldEncryptedPin := EncryptPIN(old_pin, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))
	newEncryptedPin := EncryptPIN(new_pin, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))

	body := P{
		"old_pin": oldEncryptedPin,
		"pin":     newEncryptedPin,
	}

	return MixinRequest(method, uri, body)
}

func VerifyPIN(pin string) ([]byte, error) {
	method := "POST"
	uri := "/pin/verify"
	encryptedPin := EncryptPIN(pin, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))

	body := P{
		"pin": encryptedPin,
	}

	return MixinRequest(method, uri, body)
}

func Deposit(symbol string, acount_info ...string) ([]byte, error) {
	method := "GET"
	uri := "/assets/" + symbolAssetId[symbol]
	if symbol == EOS {
		if len(acount_info) == 2 {
			body := P{
				"account_name": acount_info[0],
				"account_tag":  acount_info[1],
			}
			return MixinRequest(method, uri, body)
		}
	}

	return MixinRequest(method, uri)
}

func Withdrawal(addressId, amount, memo string) ([]byte, error) {
	method := "POST"
	uri := "/withdrawals"

	body := P{
		"address_id": addressId,
		"amount":     amount,
		"pin":        EncryptPIN(PinCode, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano())),
		"trace_id":   uuid.Must(uuid.NewV4()).String(),
		"memo":       memo,
	}

	return MixinRequest(method, uri, body)
}

func CreateAddress(symbol, address, label string, account_info ...string) ([]byte, error) {
	method := "POST"
	uri := "/addresses"

	pin := EncryptPIN(PinCode, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))

	body := P{
		"asset_id":   symbolAssetId[symbol],
		"public_key": address,
		"label":      label,
		"pin":        pin,
	}

	if symbol == EOS {
		if len(account_info) == 2 {
			body["account_name"] = account_info[0]
			body["account_tag"] = account_info[1]
		}
	}
	return MixinRequest(method, uri, body)
}

func DeleteAddress(id string) ([]byte, error) {
	method := "POST"
	uri := "/addresses/" + id + "/delete"
	pin := EncryptPIN(PinCode, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))

	body := P{
		"pin": pin,
	}
	return MixinRequest(method, uri, body)
}

func WithdrawalAddresses(symbol string) ([]byte, error) {
	method := "GET"
	uri := "/assets/" + symbolAssetId[symbol] + "/addresses"
	return MixinRequest(method, uri)
}

func ReadAddress(address_id string) ([]byte, error) {
	method := "GET"
	uri := "/addresses/" + address_id
	return MixinRequest(method, uri)
}

func ReadAsset(symbol string) ([]byte, error) {
	method := "GET"
	uri := "/assets/" + symbolAssetId[symbol]
	return MixinRequest(method, uri)
}

func ReadAssets() ([]byte, error) {
	method := "GET"
	uri := "/assets"
	return MixinRequest(method, uri)
}

func VarifyPayment(opponent_id, amount, symbol, trace_id string) ([]byte, error) {
	method := "POST"
	uri := "/payments"
	body := P{
		"asset_id":    symbolAssetId[symbol],
		"opponent_id": opponent_id,
		"amount":      amount,
		"trace_id":    trace_id,
	}
	return MixinRequest(method, uri, body)
}

func Transfer(opponent_id, amount, symbol, memo string, trace_id ...string) ([]byte, error) {
	method := "POST"
	uri := "/transfers"
	pin := EncryptPIN(PinCode, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))

	trace_uuid := uuid.Must(uuid.NewV4()).String()
	if len(trace_id) > 0 {
		trace_uuid = trace_id[0]
	}
	body := P{
		"asset_id":    symbolAssetId[symbol],
		"opponent_id": opponent_id,
		"amount":      amount,
		"pin":         pin,
		"trace_id":    trace_uuid,
		"memo":        memo,
	}

	return MixinRequest(method, uri, body)
}

func ReadTransfer(trace_id string) ([]byte, error) {
	method := "GET"
	uri := "/transfers/trace/" + trace_id
	return MixinRequest(method, uri)
}

func TopAssets() ([]byte, error) {
	method := "GET"
	uri := "/network"
	return MixinRequest(method, uri)
}

func NetworkAsset() ([]byte, error) {
	method := "GET"
	uri := "/network"
	return MixinRequest(method, uri)
}

func CreateAppUser(name string) ([]byte, string, string, error) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	block := x509.MarshalPKCS1PrivateKey(privateKey)
	b := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: block,
	}
	bt := pem.EncodeToMemory(b)

	publicKey, err := x509.MarshalPKIXPublicKey(privateKey.Public())
	if err != nil {
		return nil, "", "", err
	}

	sessionSecret := base64.StdEncoding.EncodeToString(publicKey)

	method := "POST"
	uri := "/users"

	body := P{
		"session_secret": sessionSecret,
		"full_name":      name,
	}
	resp, err := MixinRequest(method, uri, body)
	return resp, sessionSecret, string(bt), err
}
