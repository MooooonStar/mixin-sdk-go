package network

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	pin := EncryptPIN(PinCode, PinToken, SessionId, PrivateKey, uint64(time.Now().UnixNano()))
	log.Println("pin ", pin)

	method := "GET"
	uri := "/assets/" + "43d61dcd-e413-450d-80b8-101d5e903357"
	body := ""

	token, err := SignAuthenticationToken(ClientId, SessionId, PrivateKey, method, uri+ClientId, body)
	assert.Nil(t, err)
	log.Println("Bearer " + token)
}
