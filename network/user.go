package network

import (
	"encoding/json"
)

func NewUser(userId, sessionId, privateKey string, pinCodeAndToken ...string) *User {
	user := User{UserId: userId, SessionId: sessionId, PrivateKey: privateKey}
	if len(pinCodeAndToken) == 2 {
		user.SetPin(pinCodeAndToken[0], pinCodeAndToken[1])
	}
	return &user
}

func (u *User) SetPin(pinCode, pinToken string) {
	u.PinCode, u.PinToken = pinCode, pinToken
}

func (u *User) Request(method, uri string, body []byte) ([]byte, error) {
	return Request(method, uri, body, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) ReadProfile() ([]byte, error) {
	return Request("GET", "/me", nil, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) CreatePIN(oldPin, newPin string) ([]byte, error) {
	return CreatePIN(oldPin, newPin, u.PinToken, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) VerifyPIN(pin string) ([]byte, error) {
	return VerifyPIN(pin, u.PinToken, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) ReadAsset(assetID string) ([]byte, error) {
	return ReadAsset(assetID, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) ReadAssets() ([]byte, error) {
	return ReadAssets(u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) Deposit(asset string) ([]byte, error) {
	return Deposit(asset, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) CreateAddress(assetID, publicOrName, labelOrTag string) ([]byte, error) {
	return CreateAddress(assetID, publicOrName, labelOrTag, u.PinCode, u.PinToken, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) Withdrawal(assetID, publicOrName, labelOrTag, amount, memo, trace string) ([]byte, error) {
	data, err := u.CreateAddress(assetID, publicOrName, labelOrTag)
	if err != nil {
		return nil, err
	}

	type Resp struct {
		Data struct {
			AddressID string `json:"address_id"`
		} `json:"data"`
	}

	var resp Resp
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	addressID := resp.Data.AddressID
	return Withdrawal(addressID, amount, memo, trace, u.PinCode, u.PinToken, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) VerifyPayment(opponentID, amount, assetID, traceID string) ([]byte, error) {
	return VerifyPayment(opponentID, amount, assetID, traceID, u.UserId, u.SessionId, u.PrivateKey)
}

func (u User) Transfer(opponentID, amount, assetID, memo, traceID string) ([]byte, error) {
	return Transfer(opponentID, amount, assetID, memo, traceID, u.PinCode, u.PinToken, u.UserId, u.SessionId, u.PrivateKey)
}
