package main

import (
	"context"
	"testing"

	"github.com/JormungandrK/microservice-security/auth"
	"github.com/JormungandrK/microservice-user-profile/app/test"
	"github.com/JormungandrK/microservice-user-profile/db"
	"github.com/goadesign/goa"
)

var (
	service               = goa.New("user-profile-test")
	database              = db.New()
	ctrl                  = NewUserProfileController(service, database)
	hexObjectID           = "5975c461f9f8eb02aae053f3"
	internalErrorObjectID = "6975c461f9f8eb02aae053f3"
	fakeHexObjectID       = "fakeobjectidab02aae053f3"
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

func TestGetMyProfileUserProfileOK(t *testing.T) {
	// Call generated test helper, this checks that the returned media type is of the
	// correct type (i.e. uses view "default") and validates the media type.
	// Also, it ckecks the returned status code
	ctx := context.Background()
	authObj := &auth.Auth{UserID: "5975c461f9f8eb02aae053f3"}
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
	authObj := &auth.Auth{UserID: "fakeobjectidab02aae053f3"}
	ctx = auth.SetAuth(ctx, authObj)

	test.GetMyProfileUserProfileNotFound(t, ctx, service, ctrl)
}

func TestGetMyProfileUserProfileInternalServerError(t *testing.T) {
	// The test helper takes care of validating the status code for us
	ctx := context.Background()
	authObj := &auth.Auth{UserID: "6975c461f9f8eb02aae053f3"}
	ctx = auth.SetAuth(ctx, authObj)

	test.GetMyProfileUserProfileInternalServerError(t, ctx, service, ctrl)
}
