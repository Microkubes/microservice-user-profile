package main

import (
	"fmt"

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
func NewUserProfileController(service *goa.Service) *UserProfileController {
	return &UserProfileController{Controller: service.NewController("UserProfileController")}
}

// GetUserProfile runs the GetUserProfile action.
func (c *UserProfileController) GetUserProfile(ctx *app.GetUserProfileUserProfileContext) error {
	// UserProfileController_GetUserProfile: start_implement

	// Put your logic here

	// Basically you do:
	res, err := c.Repository.GetUserProfile(ctx.UserID)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	if res == nil {
		return ctx.NotFound(fmt.Errorf("User Profile not found"))
	}

	// UserProfileController_GetUserProfile: end_implement
	return ctx.OK(res)
}
