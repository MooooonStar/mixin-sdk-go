package network

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var moon_id = "825d5134-c921-3cf9-a83b-848b73"
var moon_cnb = "0x0EC770FD731C04DcDdDBca97176DC5f6af2AbeF4"

func TestVerifyPin(t *testing.T) {
	data, err := VerifyPIN("759948", PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestDeposit(t *testing.T) {
	data, err := Deposit("CNB", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestCreateAddress(t *testing.T) {
	data, err := CreateAddress("CNB", moon_cnb, "CNB Address", PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(data)
}

func TestReadAddresses(t *testing.T) {
	data, err := ReadAddress("814a0195-2048-4e09-b932-48f0b39b559b", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadAsset(t *testing.T) {
	data, err := ReadAsset("CNB", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println("cnb")
	log.Println(string(data))
}

func TestReadAssets(t *testing.T) {
	data, err := ReadAssets(UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println("hello")
	log.Println(data)
}

func TestVerifyPayment(t *testing.T) {
	data, err := VarifyPayment("825d5134-c921-3cf9-a83b-848b73c9e83b", "10", "CNB", "34fd7fee-6b14-4a24-82e1-6411768b9370", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(data)
}

func TestTransfer(t *testing.T) {
	data, err := Transfer(moon_id, "10", "CNB", "transfer test", PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(data)
}

func TestWithdraw(t *testing.T) {
	data, err := Withdrawal("814a0195-2048-4e09-b932-48f0b39b559b", "10", "from mibot", PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(data)
}

func TestReadTransfer(t *testing.T) {
	data, err := ReadTransfer("5a882c2c-6ea9-4f57-94b6-484f713d3f82", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestWithdrawalAddresses(t *testing.T) {
	data, err := WithdrawalAddresses("CNB", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(data)
}

func TestTopAssets(t *testing.T) {
	data, err := TopAssets(UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(data)
}

// func TestNetworkSnapshots(t *testing.T) {
// 	data, err := NetworkSnapshots("BTC", "", "10", "ASC")
// 	assert.Nil(t, err)
// 	log.Println(data)
// }

// func TestNetworkSnapshot(t *testing.T) {
// 	data, err := NetworkSnapshot("04ab3fe1-c817-45b3-a81d-26852b80b200")
// 	assert.Nil(t, err)
// 	log.Println(data)
// }

// func TestExternalTransactions(t *testing.T) {
// 	data, err := ExternalTransactions("CNB", moon_cnb, "", "")
// 	assert.Nil(t, err)
// 	log.Println(data)
// }

func TestSearchUser(t *testing.T) {
	data, err := MixinRequest("GET", "/search/"+"37194514", nil, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	fmt.Println("data:", string(data))
}

func TestReadUser(t *testing.T) {
	data, err := MixinRequest("GET", "/users/"+"d24fae70-32a0-453d-b5a8-980b76565297", nil, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	fmt.Println("data:", string(data))
}

var BotSessionSecret string

func TestRSA(t *testing.T) {
	fmt.Println("Save the following infomation.")

	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	block := x509.MarshalPKCS1PrivateKey(privateKey)
	privateSecret := base64.StdEncoding.EncodeToString(block)
	fmt.Println("privateSecret:", privateSecret)

	b := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: block,
	}
	bt := pem.EncodeToMemory(b)
	fmt.Println("bt", string(bt))
	assert.Nil(t, pem.Encode(os.Stdout, b))

	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(privateKey.Public())
	pubBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	assert.Nil(t, pem.Encode(os.Stdout, pubBlock))

	BotSessionSecret := base64.StdEncoding.EncodeToString(publicKeyBytes)
	fmt.Println("sessionSecret:", BotSessionSecret)

}

func TestCreateAppUser(t *testing.T) {
	user, err := CreateAppUser("coco", "123456", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	// user := User{
	// 	//ID:           "7000101596",
	// 	FullName:   "snow",
	// 	UserId:     UserId,
	// 	PinCode:    PinCode,
	// 	SessionId:  SessionId,
	// 	PinToken:   PinToken,
	// 	PrivateKey: PrivateKey,
	// }

	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	})
	bt, err := json.Marshal(user)
	assert.Nil(t, err)
	_, err = client.HSet("mixin_users", user.FullName, string(bt)).Result()
	assert.Nil(t, err)
	log.Println("")
}
