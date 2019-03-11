// GIN 认证中间件
// 时间 2019年3月9号
// 作者 申杨杨

package middleware

import (
	"github.com/gin-gonic/gin"
	"gohighfaner/pkg/gredis"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO 检查 token 是否合法
		if token := c.GetHeader("token"); token != "" {
			var id string
			if err := gredis.Get(token, &id); err == nil {
				if id == c.GetHeader("uid") {
					c.Next()
				}
			}
		}
	}
}
