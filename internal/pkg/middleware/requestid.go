package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nico612/blog-test/internal/pkg/know"
)

// RequestID 创建了一个 32 位的 UUID，并分别设置在 *gin.Context 和 HTTP 返回头中。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(know.XRequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将RequestID 保存在gin.Context 中，方便后边程序使用
		c.Set(know.XRequestIDKey, requestID)

		// 将RequestID 保存在HTTP返回头中，Header 的键为 `X-Request-ID`
		c.Writer.Header().Set(know.XRequestIDKey, requestID)

		// 在中间件中调用 Next() 方法，Next() 方法之前的代码会在到达请求方法前执行，Next() 方法之后的代码则在请求方法处理后执行
		c.Next() // 这里之后没有处理逻辑，这里只起到示例作用

		// ... 请求方法处理之后执行
	}
}
