package db

import (
	"sync"

	"github.com/Microkubes/backends"
	"github.com/Microkubes/microservice-user-profile/app"
	"github.com/keitaroinc/goa"
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
func (db *DB) GetUserProfile(objectID string) (*app.UserProfile, error) {
	db.Lock()
	defer db.Unlock()
	mediaType := &app.UserProfile{}

	if objectID == "6975c461f9f8eb02aae053f3" {
		return nil, goa.ErrInternal("Internal error")
	}

	if objectID == "fakeobjectidab02aae053f3" {
		return nil, backends.ErrNotFound("not found")
	}

	if objectID == "fakeobjectidab02aae053f3aasadas" {
		return nil, backends.ErrInvalidInput("invalid User ID")
	}

	if user, ok := db.users[objectID]; ok {
		mediaType.UserID = objectID
		mediaType.Email = &user.Email
		mediaType.FullName = &user.FullName
		mediaType.CreatedOn = 1502722729
	}

	return mediaType, nil
}

// UpdateUserProfile mock implementation
func (db *DB) UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {
	if userID == "6975c461f9f8eb02aae053f3" {
		return nil, goa.ErrInternal("Internal error")
	}

	if userID == "fakeobjectidab02aae053f3" {
		return nil, backends.ErrNotFound("not found")
	}

	if userID == "fakeobjectidab02aae053f3aasadas" {
		return nil, backends.ErrInvalidInput("invalid User ID")
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

// UpdateMyProfile mock implementation
func (db *DB) UpdateMyProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {

	if userID == "6975c461f9f8eb02aae053f3" {
		return nil, goa.ErrInternal("Internal error")
	}

	if userID == "fakeobjectidab02aae053f3" {
		return nil, backends.ErrNotFound("not found")
	}

	if userID == "fakeobjectidab02aae053f3aasadas" {
		return nil, backends.ErrInvalidInput("invalid User ID")
	}

	if _, ok := db.users[userID]; ok {
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

// Find mock implementation
func (db *DB) Find(filter *Filter, sort *Sort, page, pageSize int) (*UserProfilePage, error) {
	return nil, nil
}
