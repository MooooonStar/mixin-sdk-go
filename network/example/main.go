package main

import (
	"log"

	"github.com/MooooonStar/mixin-sdk-go/network"
)

func main() {
	user := network.NewUser(network.UserId, network.SessionId, network.PrivateKey, network.PinCode, network.PinToken)
	//user := newwork.NewUser(network.UserId, network.SessionId, network.PrivateKey)

	profile, err := user.ReadProfile()
	if err != nil {
		log.Fatal("Read profile error", err)
	}
	log.Println("profile", string(profile))

	assets, err := user.ReadAssets()
	if err != nil {
		log.Fatal("Read assets error", assets)
	}
	log.Println("assets", string(assets))
}
