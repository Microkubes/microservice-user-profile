package main

import (
	"context"
	"testing"

	"github.com/Microkubes/microservice-security/auth"
	"github.com/Microkubes/microservice-user-profile/app"
	"github.com/Microkubes/microservice-user-profile/app/test"
	"github.com/Microkubes/microservice-user-profile/db"
	"github.com/goadesign/goa"
)

var (
	service               = goa.New("user-profile-test")
	database              = db.New()
	ctrl                  = NewUserProfileController(service, database)
	hexObjectID           = "5975c461f9f8eb02aae053f3"
	internalErrorObjectID = "6975c461f9f8eb02aae053f3"
	fakeHexObjectID       = "fakeobjectidab02aae053f3"
	badHexObjectID        = "fakeobjectidab02aae053f3aasadas"
	expectedUserEmail     = "frieda@oberbrunnerkirlin.name"
	expectedUserFullName  = "Alexandra Anderson"
)

func TestGetUserProfileUserProfileOK(t *testing.T) {
	// Call generated test helper, this checks that the returned media type is of the
	// correct type (i.e. uses view "default") and validates the media type.
	// Also, it ckecks the returned status code
	_, user := test.GetUserProfileUserProfileOK(t, context.Background(), service, ctrl, hexObjectID)

	if user == nil {
		t.Fatal("Nil user")
	}

	if user.UserID != hexObjectID {
		t.Errorf("Invalid user ID, expected %s, got %s", hexObjectID, user.UserID)
	}

	userEmail := *user.Email
	if userEmail != expectedUserEmail {
		t.Errorf("Invalid user Email, expected %s, got %s", expectedUserEmail, userEmail)
	}

	userFullName := *user.FullName
	if userFullName != expectedUserFullName {
		t.Errorf("Invalid user Full Name, expected %s, got %s", expectedUserFullName, userFullName)
	}
}

func TestGetUserProfileUserProfileNotFound(t *testing.T) {
	// The test helper takes care of validating the status code for us
	test.GetUserProfileUserProfileNotFound(t, context.Background(), service, ctrl, fakeHexObjectID)
}

func TestGetUserProfileUserProfileInternalServerError(t *testing.T) {
	// The test helper takes care of validating the status code for us
	test.GetUserProfileUserProfileInternalServerError(t, context.Background(), service, ctrl, internalErrorObjectID)
}

func TestGetUserProfileUserProfileBadRequest(t *testing.T) {
	// The test helper takes care of validating the status code for us
	test.GetUserProfileUserProfileBadRequest(t, context.Background(), service, ctrl, badHexObjectID)
}

func TestGetMyProfileUserProfileOK(t *testing.T) {
	// Call generated test helper, this checks that the returned media type is of the
	// correct type (i.e. uses view "default") and validates the media type.
	// Also, it ckecks the returned status code
	ctx := context.Background()
	authObj := &auth.Auth{UserID: hexObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	_, user := test.GetMyProfileUserProfileOK(t, ctx, service, ctrl)

	if user == nil {
		t.Fatal("Nil user")
	}

	if user.UserID != hexObjectID {
		t.Errorf("Invalid user ID, expected %s, got %s", hexObjectID, user.UserID)
	}

	userEmail := *user.Email
	if userEmail != expectedUserEmail {
		t.Errorf("Invalid user Email, expected %s, got %s", expectedUserEmail, userEmail)
	}

	userFullName := *user.FullName
	if userFullName != expectedUserFullName {
		t.Errorf("Invalid user Full Name, expected %s, got %s", expectedUserFullName, userFullName)
	}

}

func TestGetMyProfileUserProfileNotFound(t *testing.T) {
	// The test helper takes care of validating the status code for us
	ctx := context.Background()
	authObj := &auth.Auth{UserID: fakeHexObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	test.GetMyProfileUserProfileNotFound(t, ctx, service, ctrl)
}

func TestGetMyProfileUserProfileInternalServerError(t *testing.T) {
	// The test helper takes care of validating the status code for us
	ctx := context.Background()
	authObj := &auth.Auth{UserID: internalErrorObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	test.GetMyProfileUserProfileInternalServerError(t, ctx, service, ctrl)
}

func TestGetMyProfileUserProfileBadRequest(t *testing.T) {
	// The test helper takes care of validating the status code for us
	ctx := context.Background()
	authObj := &auth.Auth{UserID: badHexObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	test.GetMyProfileUserProfileBadRequest(t, ctx, service, ctrl)
}

func TestUpdateUserProfileUserProfileInternalServerError(t *testing.T) {
	ctx := context.Background()

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}
	_, users := test.UpdateUserProfileUserProfileInternalServerError(t, ctx, service, ctrl, internalErrorObjectID, userProfilePayload)
	if users == nil {
		t.Fatal()
	}
}

func TestUpdateUserProfileUserProfileBadRequest(t *testing.T) {
	ctx := context.Background()

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}
	_, users := test.UpdateUserProfileUserProfileBadRequest(t, ctx, service, ctrl, badHexObjectID, userProfilePayload)
	if users == nil {
		t.Fatal()
	}
}

func TestUpdateUserProfileUserProfileOK(t *testing.T) {
	ctx := context.Background()

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}
	_, users := test.UpdateUserProfileUserProfileOK(t, ctx, service, ctrl, hexObjectID, userProfilePayload)
	if users == nil {
		t.Fatal()
	}
}

func TestUpdateMyProfileUserProfileBadRequest(t *testing.T){
	ctx := context.Background()
	authObj := &auth.Auth{UserID: badHexObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}

	_, users := test.UpdateMyProfileUserProfileBadRequest(t, ctx, service, ctrl, userProfilePayload)

	if users == nil {
		t.Fatal()
	}
}

// func UpdateMyProfileUserProfileOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserProfileController, payload *app.UserProfilePayload) (http.ResponseWriter, *app.UserProfile) {
func TestUpdateMyProfileUserProfileOK(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: hexObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}

	_, users := test.UpdateMyProfileUserProfileOK(t, ctx, service, ctrl, userProfilePayload)

	if users == nil {
		t.Fatal()
	}
}

// // func UpdateMyProfileUserProfileInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserProfileController, payload *app.UserProfilePayload) (http.ResponseWriter, error) {
func TestUpdateMyProfileUserProfileInternalServerError(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: internalErrorObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}

	_, users := test.UpdateMyProfileUserProfileInternalServerError(t, ctx, service, ctrl, userProfilePayload)

	if users == nil {
		t.Fatal()
	}}

// // func UpdateMyProfileUserProfileNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.UserProfileController, payload *app.UserProfilePayload) (http.ResponseWriter, error) {
func TestUpdateMyProfileUserProfileNotFound(t *testing.T) {
	ctx := context.Background()
	authObj := &auth.Auth{UserID: fakeHexObjectID}
	ctx = auth.SetAuth(ctx, authObj)

	userProfilePayload := &app.UserProfilePayload{
		FullName: expectedUserFullName,
		Email:    expectedUserEmail,
	}

	_, users := test.UpdateMyProfileUserProfileNotFound(t, ctx, service, ctrl, userProfilePayload)

	if users == nil {
		t.Fatal()
	}}
