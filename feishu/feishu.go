package feishu

import (
	"github.com/CatchZeng/feishu/pkg/feishu"
	"os"
)

var client *feishu.Client

func SendCard(title, body string) (err error) {
	msg := feishu.NewPostMessage()
	text := feishu.NewText(body)
	line := []feishu.PostItem{text}
	msg.SetZHTitle(title).AppendZHContent(line)
	_, _, err = client.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	token := os.Getenv("FEISHU_TOKEN")
	secret := os.Getenv("FEISHU_SECRET")
	client = feishu.NewClient(token, secret)
}
