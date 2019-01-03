package messenger

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestMsg(t *testing.T) {
	//memo := "eyJhY3Rpb24iOiJVUERBVEUiLCJwYXJ0aWNpcGFudF9pZCI6IiIsInVzZXJfaWQiOiI3YjNmMGE5NS0zZWU5LTRjMWItOGFlOS0xNzBlMzg3N2Q5MDkifQ=="
	//memo := "eyJhY3Rpb24iOiJVUERBVEUiLCJwYXJ0aWNpcGFudF9pZCI6IiIsInVzZXJfaWQiOiI3YjNmMGE5NS0zZWU5LTRjMWItOGFlOS0xNzBlMzg3N2Q5MDkifQ=="
	//memo := "eyJhY3Rpb24iOiJDUkVBVEUiLCJwYXJ0aWNpcGFudF9pZCI6IiIsInVzZXJfaWQiOiJjN2ZmNzA0ZS0xYTc0LTRmMTItYjA1Yy03YTJiZTk1NWE3ODIifQ=="
	memo := "eyJhY3Rpb24iOiJVUERBVEUiLCJwYXJ0aWNpcGFudF9pZCI6IiIsInVzZXJfaWQiOiI3YjNmMGE5NS0zZWU5LTRjMWItOGFlOS0xNzBlMzg3N2Q5MDkifQ=="
	bt, err := base64.StdEncoding.DecodeString(memo)
	if err != nil {
		panic(err)
	}
	log.Println("msg:", string(bt))
}
