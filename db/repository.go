package db

import (
	"github.com/JormungandrK/microservice-user-profile/app"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserProfileRepository defaines the interface for accessing the user profile data
type UserProfileRepository interface {
	// GetUserProfile looks up a UserProfile by the user ID.
	GetUserProfile(objectID bson.ObjectId, mediaType *app.UserProfile) error

	// UpdateUserProfile updates the UserProfile data for a particular user by its user ID.
	// If the profile already exists, it updates the data. If not profile entry exists, a new one is created.
	// Returns the updated or newly created user profile.
	//UpdateUserProfile(profile app.UserProfilePayload) (*app.UserProfile, error)
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// Find user profile by Id. return media type if success.
func (c *MongoCollection) GetUserProfile(objectID bson.ObjectId, mediaType *app.UserProfile) error {
	if err := c.FindId(objectID).One(&mediaType); err != nil {
		return err
	}

	return nil
}
