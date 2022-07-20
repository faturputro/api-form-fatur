package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send(c *gin.Context, value interface{}) {
	c.JSON(http.StatusOK, value)
}

func SendErr(c *gin.Context, value interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, value)
}
