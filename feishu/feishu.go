package feishu

import (
	"encoding/json"
	"github.com/CatchZeng/feishu/pkg/feishu"
	"log"
	"os"
)

var client *feishu.Client

type Card struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Tag  string `json:"tag"`
	Text struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	} `json:"text"`
}

func SendCard(body string) (err error) {
	element := Element{
		Tag: "div",
		Text: struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		}{
			Content: body,
			Tag:     "lark_md",
		},
	}

	card := Card{
		Elements: []Element{element},
	}

	res, err := json.Marshal(card)
	if err != nil {
		return err
	}

	msg := feishu.NewInteractiveMessage()
	msg.SetCard(string(res))
	_, _, err = client.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	token := os.Getenv("FEISHU_TOKEN")
	log.Print("token", token)
	secret := os.Getenv("FEISHU_SECRET")
	log.Print("secret", secret)
	client = feishu.NewClient(token, secret)
}
