package db

import (
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/goadesign/goa"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserProfileRepository defaines the interface for accessing the user profile data
type UserProfileRepository interface {
	// GetUserProfile looks up a UserProfile by the user ID.
	GetUserProfile(userID string, mediaType *app.UserProfile) error

	// UpdateUserProfile updates the UserProfile data for a particular user by its user ID.
	// If the profile already exists, it updates the data. If not profile entry exists, a new one is created.
	// Returns the updated or newly created user profile.
	// UpdateUserProfile(profile app.UserProfilePayload) (*app.UserProfile, error)
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// Find user profile by Id. Return media type if succeed.
func (c *MongoCollection) GetUserProfile(userID string, mediaType *app.UserProfile) error {
	// Return whether userID is a valid hex representation of object id.
	if bson.IsObjectIdHex(userID) != true {
		return goa.ErrInternal("Invalid User Id")
	}

	// Return an ObjectId from the provided hex representation.
	objectID := bson.ObjectIdHex(userID)

	// Return true if objectID is valid. A valid objectID must contain exactly 12 bytes.
	if objectID.Valid() != true {
		return goa.ErrInternal("Invalid object Id")
	}

	if err := c.FindId(objectID).One(&mediaType); err != nil {
		return goa.ErrInternal(err)
	}

	return nil
}
