package db

import (
	"sync"

	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/goadesign/goa"
)

// DB emulates a database driver using in-memory data structures.
type DB struct {
	sync.Mutex
	users map[string]*app.UserProfilePayload
}

// New initializes a new "DB" with dummy data.
func New() *DB {
	email := "frieda@oberbrunnerkirlin.name"
	fullname := "Alexandra Anderson"
	user := &app.UserProfilePayload{
		Email:    email,
		FullName: fullname,
	}
	return &DB{users: map[string]*app.UserProfilePayload{"5975c461f9f8eb02aae053f3": user}}
}

// GetUserProfile mock implementation
func (db *DB) GetUserProfile(objectID string, mediaType *app.UserProfile) error {
	db.Lock()
	defer db.Unlock()

	if objectID == "6975c461f9f8eb02aae053f3" {
		return goa.ErrInternal("Internal error")
	}

	if objectID == "fakeobjectidab02aae053f3" {
		return goa.ErrNotFound("not found")
	}

	if user, ok := db.users[objectID]; ok {
		mediaType.UserID = objectID
		mediaType.Email = &user.Email
		mediaType.FullName = &user.FullName
		mediaType.CreatedOn = 1502722729
	}

	return nil
}

func (db *DB) UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {

	if userID == "6975c461f9f8eb02aae053f3" {
		return nil, goa.ErrInternal("Internal error")
	}

	if _, ok := db.users[userID]; ok {
		db.users[userID] = &app.UserProfilePayload{
			Email:    profile.Email,
			FullName: profile.FullName,
		}
	} else {
		db.users[userID] = &app.UserProfilePayload{
			Email:    profile.Email,
			FullName: profile.FullName,
		}
	}

	res := &app.UserProfile{
		UserID:   userID,
		FullName: &profile.FullName,
		Email:    &profile.Email,
	}

	return res, nil
}
