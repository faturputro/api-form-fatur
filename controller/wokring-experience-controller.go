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

type WorkExperienceController interface {
	GetWorkExp(ctx *gin.Context)
	AddWorkExp(ctx *gin.Context)
}

type workExperienceController struct {
	s services.WorkExperienceService
}

func NewWorkExperienceController(s services.WorkExperienceService) WorkExperienceController {
	return &workExperienceController{
		s: s,
	}
}
func (pc *workExperienceController) GetWorkExp(ctx *gin.Context) {
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
	desc, err := pc.s.Get(uint(id))
	res = helpers.BuildResponse(true, "", dto.WorkExperience{
		Description: desc,
	})
	ctx.JSON(http.StatusCreated, res)
}

func (pc *workExperienceController) AddWorkExp(ctx *gin.Context) {
	var (
		req entities.WorkExperience
		res helpers.Response
	)
	err := ctx.Bind(&req)
	id, err := strconv.Atoi(ctx.Param("id"))
	we, err := pc.s.Add(req, uint(id))
	if err != nil {
		res = helpers.BuildErrorResponse("Something went wrong", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res = helpers.BuildResponse(true, "Profile successfully created.", dto.WorkExperience{
		Description: we,
	})
	ctx.JSON(http.StatusCreated, res)
}
