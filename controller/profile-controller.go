package controller

import (
	"dummy/dto"
	"dummy/entities"
	"dummy/helpers"
	"dummy/services"
	"net/http"
	"strconv"

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

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		res = helpers.BuildErrorResponse("Failed to process request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	profile, err := pc.s.Get(uint(id))
	res = helpers.BuildResponse(true, "", profile)
	ctx.JSON(http.StatusCreated, res)
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
	var (
		req entities.Profile
		res helpers.Response
	)
	id, err := strconv.Atoi(ctx.Param("id"))
	err = ctx.Bind(&req)
	profileId, err := pc.s.Update(req, uint(id))
	if err != nil {
		res = helpers.BuildErrorResponse("Something went wrong", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res = helpers.BuildResponse(true, "Profile successfully updated.", dto.Profile{
		ProfileCode: profileId,
	})
	ctx.JSON(http.StatusCreated, res)
}
