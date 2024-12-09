package sms

import (
	"feishu-gin/feishu"
	"github.com/gin-gonic/gin"
	"log"
)

type MsgBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type returnMsg map[string]interface{}

func SendSms(r *gin.Context) {
	msgBody := &MsgBody{}
	if err := r.ShouldBindJSON(msgBody); err != nil {
		r.JSON(200, returnMsg{"code": 400, "msg": err})
		log.Print("参数解析失败", err)
		return
	}
	if err := feishu.SendCard(msgBody.Title, msgBody.Body); err != nil {
		r.JSON(200, returnMsg{"code": 400, "msg": err})
		log.Print("发送失败", err)
		return
	}
	r.JSON(200, returnMsg{"code": 200, "msg": "发送成功"})
}
