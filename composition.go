package main

import (
	"dummy/cache"
	"dummy/controller"
	"dummy/repository"
	"dummy/services"

	"gorm.io/gorm"
)

type Composition struct {
	profileRepo repository.ProfileRepo
	profile     controller.ProfileController
	workExp     controller.WorkExperienceController
}

func ControllerInstance(db *gorm.DB, redis cache.AppCache) Composition {
	var (
		profileRepo              repository.ProfileRepo              = repository.NewProfileRepo(db)
		profileService           services.ProfileService             = services.NewProfileService(redis, profileRepo)
		profileController        controller.ProfileController        = controller.NewProfileController(profileService)
		workExpRepo              repository.WorkExperienceRepository = repository.NewWorkExperienceRepository(db)
		workExpService           services.WorkExperienceService      = services.NewWorkExperienceService(workExpRepo)
		workExperienceController controller.WorkExperienceController = controller.NewWorkExperienceController(workExpService)
	)

	return Composition{
		profileRepo: profileRepo,
		profile:     profileController,
		workExp:     workExperienceController,
	}
}
