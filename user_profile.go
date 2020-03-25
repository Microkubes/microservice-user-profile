package main

import (
	"fmt"

	"github.com/Microkubes/backends"
	errors "github.com/Microkubes/backends"
	"github.com/Microkubes/microservice-security/auth"
	"github.com/Microkubes/microservice-user-profile/app"
	"github.com/Microkubes/microservice-user-profile/db"
	"github.com/keitaroinc/goa"
)

const MaxResultsPerPage = 500

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
	res, err := c.Repository.GetUserProfile(ctx.UserID)

	if err != nil {
		if errors.IsErrNotFound(err) {
			return ctx.NotFound(err)
		}

		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(err)
		}

		return ctx.InternalServerError(err)
	}

	res.UserID = ctx.UserID

	return ctx.OK(res)
}

// UpdateUserProfile runs the UpdateUserProfile action.
func (c *UserProfileController) UpdateUserProfile(ctx *app.UpdateUserProfileUserProfileContext) error {
	res, err := c.Repository.UpdateUserProfile(ctx.Payload, ctx.UserID)

	if err != nil {
		fmt.Printf("  => ERROR:%s\n", err)

		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(err)
		}

		if errors.IsErrAlreadyExists(err) {
			return ctx.BadRequest(err)
		}

		return ctx.InternalServerError(err)
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

	fmt.Println(userID)

	res, err := c.Repository.UpdateUserProfile(ctx.Payload, userID)

	if err != nil {
		fmt.Printf("  => ERROR:%s\n", err)

		if errors.IsErrNotFound(err) {
			return ctx.NotFound(err)
		}

		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(err)
		}

		if errors.IsErrAlreadyExists(err) {
			return ctx.BadRequest(err)
		}

		return ctx.InternalServerError(err)
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
	res, err := c.Repository.GetUserProfile(userID)

	if err != nil {
		if errors.IsErrNotFound(err) {
			return ctx.NotFound(err)
		}

		if errors.IsErrInvalidInput(err) {
			return ctx.BadRequest(err)
		}

		return ctx.InternalServerError(err)
	}

	res.UserID = userID

	return ctx.OK(res)
}

// FindUserProfile runs FindUserProfile action
func (c *UserProfileController) FindUserProfile(ctx *app.FindUserProfileUserProfileContext) error {

	hasAuth := auth.HasAuth(ctx)

	if !hasAuth {
		return ctx.InternalServerError(goa.ErrUnauthorized("Auth has not been set"))
	}

	page := ctx.Payload.Page
	if page < 1 {
		return ctx.BadRequest(goa.ErrBadRequest("invalid page number"))
	}
	page--
	pageSize := ctx.Payload.PageSize
	if pageSize > MaxResultsPerPage {
		pageSize = MaxResultsPerPage
	}

	var filter *db.Filter
	filter = &db.Filter{
		Match: map[string]interface{}{},
	}
	if ctx.Payload.Filter != nil && len(ctx.Payload.Filter) > 0 {
		for _, filterProp := range ctx.Payload.Filter {
			filter.Match[filterProp.Property] = filterProp.Value
		}
	}
	var sort *db.Sort
	if ctx.Payload.Sort != nil {
		sort = &db.Sort{
			SortBy:    ctx.Payload.Sort.Property,
			Direction: ctx.Payload.Sort.Direction,
		}
	}

	result, err := c.Repository.Find(filter, sort, page, pageSize)
	if err != nil {
		if backends.IsErrInvalidInput(err) {
			return ctx.BadRequest(goa.ErrBadRequest(err))
		}
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	userProfilesPage := &app.UserProfilePage{
		Page:     &result.Page,
		PageSize: &result.PageSize,
		Items:    []*app.UserProfile{},
	}

	for _, userProfileItem := range result.Items {
		userProfile := userProfileItem

		userProfilesPage.Items = append(userProfilesPage.Items, &app.UserProfile{
			UserID:                    userProfile.UserID,
			FullName:                  &userProfile.FullName,
			Email:                     &userProfile.Email,
			Company:                   &userProfile.Company,
			TaxNumber:                 &userProfile.TaxNumber,
			CompanyRegistrationNumber: &userProfile.CompanyRegistrationNumber,
		})
	}

	return ctx.OK(userProfilesPage)
}
