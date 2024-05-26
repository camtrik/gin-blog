package middleaware

import (
	"fmt"
	"time"

	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/pkg/app"
	"github.com/camtrik/gin-blog/pkg/email"
	"github.com/camtrik/gin-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrame().Errorf(c, "panic recover err: %v", err)

				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("panic, time: %d", time.Now().Unix()),
					fmt.Sprintf("panic recover err: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
