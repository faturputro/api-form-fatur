package helpers

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateRedisKey(c *gin.Context) string {
	var key string
	url := strings.Split(c.Request.URL.String(), "/")
	split := strings.Join(url[2:], "/")
	key = fmt.Sprintf("/%s", strings.ReplaceAll(split, "?", "/query:"))

	return key
}
