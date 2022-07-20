package main

import (
	"dummy/cache"
	"dummy/config"
	"dummy/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func routes(isDebug bool) *gin.Engine {
	host := viper.GetString("REDIS_HOST")
	port := viper.GetInt("REDIS_PORT")
	rdb := viper.GetInt("REDIS_DB")
	redisPass := viper.GetString("REDIS_PASS")

	var engine *gin.Engine
	var (
		db    *gorm.DB       = config.SetupConnection()
		redis cache.AppCache = cache.NewRedisCache(host, port, rdb, redisPass)
		comp  Composition    = ControllerInstance(db, redis)
	)

	if isDebug {
		engine = gin.Default()
	} else {
		engine = gin.New()
	}
	engine.Use(middleware.CORSMiddleware())

	routes := engine.Group("/api")
	routes.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG",
		})
	})

	profile := routes.Group("/v1/profile")
	{
		profile.POST("/", comp.profile.CreateProfile)
		profile.GET("/:id", comp.profile.GetProfile)
		profile.PUT("/:id", comp.profile.UpdateProfile)
	}
	workingExperience := routes.Group("/v1/working-experience")
	{
		workingExperience.Use(middleware.ProfileValidator(comp.profileRepo, redis))
		workingExperience.GET("/:id", comp.workExp.GetWorkExp)
		workingExperience.PUT("/:id", comp.workExp.AddWorkExp)
	}
	return engine
}
