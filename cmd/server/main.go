package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Coords     []string `json:"coords"`
	Desc       string   `json:"desc"`
	Nick       string   `json:"nick"`
	TargetNick string   `json:"target_nick"`
	Wpbot      bool     `json:"wpbot"`
}

func main() {
	p, err := json.Marshal(Payload{
		Coords:     nil,
		Desc:       "",
		Nick:       "",
		TargetNick: "",
		Wpbot:      true})
	bodyReader := bytes.NewReader(p)
	if err != nil {
		log.Println("failed to marshal payload", err)
	}
	client := http.Client{}
	_, err = client.Post("https://go-pjatk-server.fly.dev/api/game", "application/json", bodyReader)
	if err != nil {
		log.Println("failed to post", err)
	}
}
