package main

import (
	"dummy/cache"
	"dummy/controller"
	"dummy/repository"
	"dummy/services"

	"gorm.io/gorm"
)

type Composition struct {
	profile controller.ProfileController
}

func ControllerInstance(db *gorm.DB, redis cache.AppCache) Composition {
	var (
		profileRepo       repository.ProfileRepo       = repository.NewProfileRepo(db)
		profileService    services.ProfileService      = services.NewProfileService(redis, profileRepo)
		profileController controller.ProfileController = controller.NewProfileController(profileService)
	)

	return Composition{
		profile: profileController,
	}
}
