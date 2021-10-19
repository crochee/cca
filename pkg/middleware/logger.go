package middleware

import (
	"github.com/crochee/lib/log"
	"github.com/gin-gonic/gin"

	"cca/pkg/v"
)

// RequestLogger 设置请求日志
func RequestLogger(logger log.Builder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fieldList []log.Field
		ctx := c.Request.Context()
		if traceID := v.GetTraceID(ctx); traceID != "" {
			fieldList = append(fieldList, log.Field{
				Key:   "TraceID",
				Value: traceID,
			})
		}
		fieldList = append(fieldList, log.Field{
			Key:   "ClientIP",
			Value: c.ClientIP(),
		})
		c.Request = c.Request.WithContext(log.WithContext(ctx, logger.With(fieldList...)))
		c.Next()
	}
}