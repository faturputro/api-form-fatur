package middleware

import (
	"dummy/cache"
	"dummy/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CacheMiddleware(client cache.AppCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := helpers.GenerateRedisKey(c)
		if c.Request.Method == "GET" {
			if !client.Exist(key) {
				var val interface{}
				_, err := client.Get(key, &val)
				if err != nil {
					fmt.Println(err)
				}
				res := helpers.BuildResponse(true, "", val)
				c.AbortWithStatusJSON(http.StatusOK, res)
				return
			} else {
				c.Next()
			}
		}
	}
}
