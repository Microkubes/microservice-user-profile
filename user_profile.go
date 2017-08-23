package main

import (
	"github.com/JormungandrK/microservice-security/auth"
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/JormungandrK/microservice-user-profile/db"
	"github.com/goadesign/goa"
)
//   "gopkg.in/mgo.v2"
// UserProfileController implements the userProfile resource.
type UserProfileController struct {
	*goa.Controller
	Repository db.UserProfileRepository
}

type ChangeInfo struct {
    Updated    int         // Number of existing documents updated
    Removed    int         // Number of documents removed
    UpsertedId interface{} // Upserted _id field, when not explicitly provided
}

// Collection is an interface to access to the collection struct.
type Collection interface {
        UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error)
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

func (c *UserProfileController) UpdateUserProfile(ctx *app.UpdateUserProfileUserProfileContext) error {
	payload := *ctx.Payload
	_, err := c.Repository.UpdateUserProfile(payload)
	if err != nil {
		return err
	}
	return nil

}

func (c *UserProfileController) GetMyProfile(ctx *app.GetMyProfileUserProfileContext) error {
	var authObj *auth.Auth

	hasAuth := auth.HasAuth(ctx)

	if hasAuth {
		authObj = auth.GetAuth(ctx.Context)
	} else {
		return ctx.InternalServerError(goa.ErrInternal("Auth has not been set"))
	}

	userID := authObj.UserID
	res := &app.UserProfile{}

	// Return one user profile by id.
	if err := c.Repository.GetUserProfile(userID, res); err != nil {
		return ctx.InternalServerError(err)
	}
	if res.CreatedOn == 0 {
		return ctx.NotFound(goa.ErrNotFound("User Profile not found"))
	}

	res.UserID = userID

	return ctx.OK(res)
}
