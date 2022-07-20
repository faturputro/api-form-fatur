package controller

import (
	"dummy/dto"
	"dummy/entities"
	"dummy/helpers"
	"dummy/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	GetProfile(ctx *gin.Context)
	CreateProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type profileController struct {
	s services.ProfileService
}

func NewProfileController(s services.ProfileService) ProfileController {
	return &profileController{
		s: s,
	}
}

func (pc *profileController) GetProfile(ctx *gin.Context) {
	var (
		res helpers.Response
		err error
	)
	if err != nil {
		res = helpers.BuildErrorResponse("Failed to process request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
}

func (pc *profileController) CreateProfile(ctx *gin.Context) {
	var (
		req entities.Profile
		res helpers.Response
	)
	err := ctx.Bind(&req)
	profileId, err := pc.s.Create(req)
	if err != nil {
		res = helpers.BuildErrorResponse("Something went wrong", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res = helpers.BuildResponse(true, "Profile successfully created.", dto.Profile{
		ProfileCode: profileId,
	})
	ctx.JSON(http.StatusCreated, res)
}
func (pc *profileController) UpdateProfile(ctx *gin.Context) {

}
