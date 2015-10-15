package main

import (
	"fmt"
	"github.com/zenazn/goji"
	"net"

	"github.com/bluele/slack"
	"github.com/goji/httpauth"

	"github.com/BurntSushi/toml"
	"github.com/ccding/go-stun/stun"
	_ "github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os/exec"
	"time"
)

type tomlConfig struct {
	Slack  slackConfig
	Server serverConfig
}

type slackConfig struct {
	Channel string
	Token   string
}

type serverConfig struct {
	TargetDomain string
	NameServer   string
}

func main() {

	go work()
	goji.Use(httpauth.SimpleBasicAuth("a", "a"))
	goji.Get("/v1/log", LogAPIContoller)
	goji.Get("/v1/stat", StatAPIContoller)
	goji.Serve()

}

func work() {
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	targetDomain := config.Server.TargetDomain
	nameServer := config.Server.NameServer

	// counter
	sameCount := 0
	checkCount := 0

	// connection slack-api
	api := slack.New(config.Slack.Token)
	channel, err := api.FindChannelByName(config.Slack.Channel)
	if err != nil {
		panic(err)
	}

	for {
		checkCount = checkCount + 1

		// notify running?
		if checkCount%60 == 0 {
			err = api.ChatPostMessage(channel.Id, "@takoyaki I'm running. ", nil)
			if err != nil {
				fmt.Println("slack api error", err)
			}
		}

		time.Sleep(60 * time.Second)

		// nonreq-resovle using dig(/hack/resolve.sh), get registry ipv4 address
		ripaddr, err := resolve(targetDomain, nameServer)
		if err != nil {
			fmt.Println("resolve error:", err)
			continue
		}

		fmt.Println("resovle ipaddr:", ripaddr.String())

		// get global-ipv4-address using STUN
		_, host, err := stun.NewClient().Discover()
		if err != nil {
			fmt.Println("stun client error", err)
			continue
		}
		gipaddr := net.ParseIP(host.IP())
		fmt.Println("stun client ipaddr:", gipaddr.String())

		// compare
		if gipaddr.String() == ripaddr.String() {
			fmt.Println("same")
			sameCount = sameCount + 1
			continue
		}

		// notify update
		err = api.ChatPostMessage(channel.Id, "@takoyaki global-ipv4-address != resolve-ipv4-address. Now, I try to update. ", nil)
		if err != nil {
			fmt.Println("slack api error", err)
		}
		fmt.Println("not equal ipaddr")

		// update
		_, err = exec.Command("/bin/bash", "hack/update.sh", gipaddr.String()).Output()
		if err != nil {
			fmt.Println("update error", err)
			continue
		}

		// wait
		time.Sleep(60 * 10 * time.Second)

	}
}
