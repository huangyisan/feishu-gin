package routers

import (
	"feishu-gin/routers/sms"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	sms.Send(apiGroup)
}
