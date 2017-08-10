package main

import (
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/goadesign/goa"
)

// UserProfileController implements the userProfile resource.
type UserProfileController struct {
	*goa.Controller
}

// NewUserProfileController creates a userProfile controller.
func NewUserProfileController(service *goa.Service) *UserProfileController {
	return &UserProfileController{Controller: service.NewController("UserProfileController")}
}

// GetUserProfile runs the GetUserProfile action.
func (c *UserProfileController) GetUserProfile(ctx *app.GetUserProfileUserProfileContext) error {
	// UserProfileController_GetUserProfile: start_implement

	// Put your logic here

	// UserProfileController_GetUserProfile: end_implement
	res := &app.UserProfile{}
	return ctx.OK(res)
}
