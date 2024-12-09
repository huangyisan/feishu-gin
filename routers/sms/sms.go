package sms

import (
	"feishu-gin/controller/sms"
	"github.com/gin-gonic/gin"
)

func Send(group *gin.RouterGroup) {
	group.POST("/send", sms.SendSms)
}
