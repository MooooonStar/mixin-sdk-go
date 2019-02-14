package network

import (
	"fmt"
	"log"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var snow = "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
var snowCNBAddr = "0x4fE05eBB326f52A671247d693a56771e29E1b5EA"

func TestVerifyPin(t *testing.T) {
	data, err := VerifyPIN(PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestDeposit(t *testing.T) {
	data, err := Deposit(CNB, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestCreateAddress(t *testing.T) {
	// data, err := CreateAddress(CNB, snowCNBAddr, "CNB Address", PinCode, PinToken, UserId, SessionId, PrivateKey)
	// assert.Nil(t, err)
	// log.Println(string(data))
	data, err := CreateAddress(EOS, "eoswithmixin", "a282d3c9e6f121db99f728a5f8e3ff64", PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadAddresses(t *testing.T) {
	data, err := ReadAddress("814a0195-2048-4e09-b932-48f0b39b559b", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadAsset(t *testing.T) {
	data, err := ReadAsset(EOS, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadAssets(t *testing.T) {
	data, err := ReadAssets(UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestVerifyPayment(t *testing.T) {
	data, err := VarifyPayment("825d5134-c921-3cf9-a83b-848b73c9e83b", "10", "CNB", "34fd7fee-6b14-4a24-82e1-6411768b9370", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestTransfer(t *testing.T) {
	trace := uuid.Must(uuid.NewV4()).String()
	data, err := Transfer(snow, "10", CNB, "transfer test", trace, PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestWithdraw(t *testing.T) {
	trace := uuid.Must(uuid.NewV4()).String()
	//data, err := Withdrawal("5dfe3f1e-7022-4f37-901d-49febaf485bf", "11", "Hello", trace, PinCode, PinToken, UserId, SessionId, PrivateKey)
	data, err := Withdrawal("4ceab4e8-79e9-4be5-8c5d-93e264ec3589", "0.0001", "Hi", trace, PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadTransfer(t *testing.T) {
	data, err := ReadTransfer("6ac2ee21-a9ef-4b52-8774-d4d18a622161", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestWithdrawalAddresses(t *testing.T) {
	data, err := WithdrawalAddresses(CNB, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestTopAssets(t *testing.T) {
	data, err := TopAssets(UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestNetworkSnapshots(t *testing.T) {
	data, err := NetworkSnapshots(XIN, time.Now().Add(-1*time.Hour), true, 3, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestNetworkSnapshot(t *testing.T) {
	data, err := NetworkSnapshot("cb7f1f3b-8987-4712-8235-e801f1ccd042", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestExternalTransactions(t *testing.T) {
	data, err := ExternalTransactions(CNB, "0x4fE05eBB326f52A671247d693a56771e29E1b5EA", "", time.Now().Add(-24*time.Hour), 10, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

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

func TestCreateAppUser(t *testing.T) {
	user, err := CreateAppUser("no one", "123456", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)

	info, err := user.ReadProfile()
	assert.Nil(t, err)
	log.Println(string(info))
}
