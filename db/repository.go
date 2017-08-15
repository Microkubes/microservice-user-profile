package db

import (
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/goadesign/goa"

	"time"
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

// NewSession returns a new Mongo Session.
func NewSession(Host string, Username string, Password string, Database string) *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{Host},
		Username: Username,
		Password: Password,
		Database: Database,
		Timeout:  30 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	// SetMode - consistency mode for the session.
	session.SetMode(mgo.Monotonic, true)

	return session
}

// PrepareDB ensure presence of persistent and immutable data in the DB.
func PrepareDB(session *mgo.Session, db string, dbCollection string, indexes []string) *mgo.Collection {
	// Create collection
	collection := session.DB(db).C(dbCollection)

	// Define indexes
	for _, elem := range indexes {
		i := []string{elem}
		index := mgo.Index{
			Key:        i,
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		// Create indexes
		if err := collection.EnsureIndex(index); err != nil {
			panic(err)
		}
	}

	return collection
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