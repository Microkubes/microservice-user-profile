package main

import (
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/JormungandrK/microservice-user-profile/db"
	"github.com/goadesign/goa"
)

// UserProfileController implements the userProfile resource.
type UserProfileController struct {
	*goa.Controller
	Repository db.UserProfileRepository
}

// NewUserProfileController creates a userProfile controller.
func NewUserProfileController(service *goa.Service, Repository db.UserProfileRepository) *UserProfileController {
	return &UserProfileController{
		Controller: service.NewController("UserProfileController"),
		Repository: Repository,
	}
}

// GetUserProfile runs the GetUserProfile action.
func (c *UserProfileController) GetUserProfile(ctx *app.GetUserProfileUserProfileContext) error {
	// Build the resource using the generated data structure.
	res := &app.UserProfile{}

	// Return one user profile by id.
	if err := c.Repository.GetUserProfile(ctx.UserID, res); err != nil {
		return ctx.InternalServerError(err)
	}
	if res.CreatedOn == 0 {
		return ctx.NotFound(goa.ErrNotFound("User Profile not found"))
	}

	res.UserID = ctx.UserID

	return ctx.OK(res)
}

func (c *UserProfileController) GetMyProfile(ctx *app.GetMyProfileUserProfileContext) error {
	// Build the resource using the generated data structure.
	res := &app.UserProfile{}

	// Return one user profile by id.
	if err := c.Repository.GetUserProfile(ctx.UserID, res); err != nil {
		return ctx.InternalServerError(err)
	}
	if res.CreatedOn == 0 {
		return ctx.NotFound(goa.ErrNotFound("User Profile not found"))
	}

	res.UserID = ctx.UserID

	return ctx.OK(res)	
}
