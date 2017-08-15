package db

import (
	"errors"
	"sync"

	"github.com/JormungandrK/microservice-user-profile/app"

	"gopkg.in/mgo.v2/bson"
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
	userId := "5975c461f9f8eb02aae053f3"
	user := &app.UserProfilePayload{
		Email:    &email,
		FullName: &fullname,
		UserID:   &userId,
	}
	return &DB{users: map[string]*app.UserProfilePayload{"5975c461f9f8eb02aae053f3": user}}
}

// GetUserProfile mock implementation
func (db *DB) GetUserProfile(objectID bson.ObjectId, mediaType *app.UserProfile) error {
	db.Lock()
	defer db.Unlock()

	id := objectID.Hex()

	if id == "6975c461f9f8eb02aae053f3" {
		err := errors.New("Internal error")
		return err
	}

	if user, ok := db.users[id]; ok {
		mediaType.UserID = id
		mediaType.Email = user.Email
		mediaType.FullName = user.FullName
		mediaType.CreatedOn = 1502722729
	}

	return nil
}
