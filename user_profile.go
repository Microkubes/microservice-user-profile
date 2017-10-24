package main

import (
	"github.com/JormungandrK/microservice-security/auth"
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/JormungandrK/microservice-user-profile/db"
	"github.com/goadesign/goa"
)

// UserProfileController implements the userProfile resource.
type UserProfileController struct {
	*goa.Controller
	Repository db.UserProfileRepository
}

// ChangeInfo keeps the number of changed documents
type ChangeInfo struct {
	Updated    int         // Number of existing documents updated
	Removed    int         // Number of documents removed
	UpsertedID interface{} // Upserted _id field, when not explicitly provided
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
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 400:
			return ctx.BadRequest(err)
		case 404:
			return ctx.NotFound(err)
		default:
			return ctx.InternalServerError(err)
		}
	}

	res.UserID = ctx.UserID

	return ctx.OK(res)
}

// UpdateUserProfile runs the UpdateUserProfile action.
func (c *UserProfileController) UpdateUserProfile(ctx *app.UpdateUserProfileUserProfileContext) error {
	res, err := c.Repository.UpdateUserProfile(ctx.Payload, ctx.UserID)

	if err != nil {
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 400:
			return ctx.BadRequest(err)
		default:
			return ctx.InternalServerError(err)
		}
	}

	return ctx.OK(res)
}

// UpdateMyProfile runs the UpdateMyProfile action.
func (c *UserProfileController) UpdateMyProfile(ctx *app.UpdateMyProfileUserProfileContext) error {
	var authObj *auth.Auth

	hasAuth := auth.HasAuth(ctx)

	if hasAuth {
		authObj = auth.GetAuth(ctx.Context)
	} else {
		return ctx.InternalServerError(goa.ErrInternal("Auth has not been set"))
	}

	userID := authObj.UserID

	res, err := c.Repository.UpdateMyProfile(ctx.Payload, userID)

	if err != nil {
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 400:
			return ctx.BadRequest(err)
		case 404:
			return ctx.NotFound(err)
		default:
			return ctx.InternalServerError(err)
		}
	}

	return ctx.OK(res)
}

// GetMyProfile runs the GetMyProfile action.
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
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 400:
			return ctx.BadRequest(err)
		case 404:
			return ctx.NotFound(err)
		default:
			return ctx.InternalServerError(err)
		}
	}

	res.UserID = userID

	return ctx.OK(res)
}
