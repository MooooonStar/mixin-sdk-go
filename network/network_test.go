package network

import (
	"fmt"
	"log"
	"testing"
	"time"

	prettyjson "github.com/hokaccha/go-prettyjson"
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
	data, err := CreateAddress(CNB, "0x4fE05eBB326f52A671247d693a56771e29E1b5EA", "haha", PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
	// data, err := CreateAddress(EOS, "eoswithmixin", "a282d3c9e6f121db99f728a5f8e3ff64", PinCode, PinToken, UserId, SessionId, PrivateKey)
	// assert.Nil(t, err)
	// log.Println(string(data))
}

func TestReadAddresses(t *testing.T) {
	data, err := ReadAddress("814a0195-2048-4e09-b932-48f0b39b559b", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadAsset(t *testing.T) {
	data, err := ReadAsset(BTC, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestReadAssets(t *testing.T) {
	data, err := ReadAssets(UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestVerifyPayment(t *testing.T) {
	data, err := VerifyPayment("825d5134-c921-3cf9-a83b-848b73c9e83b", "10", "CNB", "34fd7fee-6b14-4a24-82e1-6411768b9370", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestTransfer(t *testing.T) {
	trace := uuid.Must(uuid.NewV4()).String()
	data, err := Transfer("7b3f0a95-3ee9-4c1b-8ae9-170e3877d909", "10", CNB, "test transfer", trace, PinCode, PinToken, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	v, _ := prettyjson.Format(data)
	log.Println(string(v))
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

func TestNetworkAssets(t *testing.T) {
	data, err := NetworkAsset(BTC, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestNetworkSnapshots(t *testing.T) {
	// 2019-04-08T07:34:16.556276Z
	checkpoint, _ := time.Parse(time.RFC3339Nano, "2019-04-08T05:33:41.100000Z")
	data, err := NetworkSnapshots("", checkpoint, "DESC", 10, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	v, _ := prettyjson.Format(data)
	log.Println(string(v))
}

func TestNetworkSnapshot(t *testing.T) {
	data, err := NetworkSnapshot("c95108e9-81e7-4119-93bd-1674ed121bbf", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestExternalTransactions(t *testing.T) {
	offset, _ := time.Parse(time.RFC3339Nano, "2019-01-01T15:04:05.999999999Z")
	data, err := ExternalTransactions("", "", "", offset, 10, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	log.Println(string(data))
}

func TestSearchUser(t *testing.T) {
	data, err := MixinRequest("GET", "/search/37066011", nil, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	fmt.Println("data:", string(data))
}

func TestReadUser(t *testing.T) {
	data, err := MixinRequest("GET", "/users/"+"14521f6b-2619-41ba-89ff-d440330cbde0", nil, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	fmt.Println("data:", string(data))
}

func TestReadProfile(t *testing.T) {
	data, err := MixinRequest("GET", "/me", nil, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	fmt.Println("data:", string(data))
}

func TestUpdateProfile(t *testing.T) {
	params := P{
		"redirect_uri": "http://www.igroup.pub/auth",
	}
	data, err := MixinRequest("POST", "/me", params, UserId, SessionId, PrivateKey)
	assert.Nil(t, err)
	fmt.Println("data:", string(data))
}

func TestCreateAppUser(t *testing.T) {
	user, err := CreateAppUser("no one", "123456", UserId, SessionId, PrivateKey)
	assert.Nil(t, err)

	v0, _ := prettyjson.Marshal(user)
	log.Println(string(v0))

	info, err := user.ReadProfile()
	assert.Nil(t, err)
	log.Println(string(info))
}
