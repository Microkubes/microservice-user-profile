package main

import (
	"context"
	"testing"

	"github.com/JormungandrK/microservice-user-profile/app/test"
	"github.com/JormungandrK/microservice-user-profile/db"
	"github.com/goadesign/goa"
)

var (
	service         = goa.New("user-profile-test")
	database        = db.New()
	ctrl            = NewUserProfileController(service, database)
	HexObjectID     = "5975c461f9f8eb02aae053f3"
	DiffHexObjectID = "6975c461f9f8eb02aae053f3"
	FakeHexObjectID = "fakeobjectidab02aae053f3"
)

func GetUserProfileUserProfileOK(t *testing.T) {
	// Call generated test helper, this checks that the returned media type is of the
	// correct type (i.e. uses view "default") and validates the media type.
	// Also, it ckecks the returned status code
	_, user := test.GetUserOK(t, context.Background(), service, ctrl, HexObjectID)

	if user == nil {
		t.Fatal("Nil user")
	}

	if user.ID != HexObjectID {
		t.Errorf("Invalid user ID, expected %s, got %s", HexObjectID, user.ID)
	}
}
