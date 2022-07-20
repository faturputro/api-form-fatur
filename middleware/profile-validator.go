package middleware

import (
	"dummy/cache"
	"dummy/helpers"
	"dummy/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ProfileValidator(pr repository.ProfileRepo, client cache.AppCache) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var res helpers.Response
		id, err := strconv.Atoi(ctx.Param("id"))
		found, err := pr.FindById(uint(id))
		if err != nil {
			res = helpers.BuildErrorResponse("Failed to process request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		if found.ID == 0 {
			err = errors.New("Profile not found")
			res = helpers.BuildErrorResponse("Failed to process request", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}
		ctx.Next()
	}
}
