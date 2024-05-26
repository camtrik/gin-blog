package middleaware

import (
	"bytes"
	"time"

	"github.com/camtrik/gin-blog/global"
	"github.com/camtrik/gin-blog/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next() // execute next middleware or request, log time
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(), // request data after encoding
			"response": bodyWriter.body.String(),    // response data
		}

		global.Logger.WithFields(fields).Infof(c, "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method, bodyWriter.Status(), beginTime, endTime,
		)
	}
}
