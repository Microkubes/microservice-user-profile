package db

import (
	"errors"
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
	userId := "5975c461f9f8eb02aae053f3"
	user := &app.UserProfilePayload{
		Email:    &email,
		FullName: &fullname,
		UserID:   &userId,
	}
	return &DB{users: map[string]*app.UserProfilePayload{"5975c461f9f8eb02aae053f3": user}}
}

// GetUserProfile mock implementation
func (db *DB) GetUserProfile(objectID string, mediaType *app.UserProfile) error {
	db.Lock()
	defer db.Unlock()

	if objectID == "6975c461f9f8eb02aae053f3" {
		err := errors.New("Internal error")
		return goa.ErrInternal(err)
	}

	if user, ok := db.users[objectID]; ok {
		mediaType.UserID = objectID
		mediaType.Email = user.Email
		mediaType.FullName = user.FullName
		mediaType.CreatedOn = 1502722729
	}

	return nil
}

func (db *DB) UpdateUserProfile(profile app.UserProfilePayload) (*app.UserProfile, error){

	if *profile.UserID == "6975c461f9f8eb02aae053f3" {
		err := errors.New("Internal error")
		return nil, goa.ErrInternal(err)
	}


    if _, ok := db.users[*profile.UserID]; ok {
        db.users[*profile.UserID] =  &app.UserProfilePayload{
            Email:    profile.Email,
            FullName: profile.FullName,
            UserID:   profile.UserID,
			CreateOn: profile.CreateOn,
        }
    }else{
        db.users[*profile.UserID] = &app.UserProfilePayload{
            Email:    profile.Email,
            FullName: profile.FullName,
            UserID:   profile.UserID,
			CreateOn: profile.CreateOn,
        }
    }

	res := &app.UserProfile{
		UserID:     *profile.UserID,
		FullName:   profile.FullName,
		Email:      profile.Email,
	}

    // return db.users[*profile.UserID], nil

	return res, nil
}