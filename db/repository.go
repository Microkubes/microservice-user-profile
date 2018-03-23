package db

import (
	"github.com/Microkubes/microservice-user-profile/app"
	"github.com/goadesign/goa"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// UserProfileRepository defaines the interface for accessing the user profile data
type UserProfileRepository interface {
	// GetUserProfile looks up a UserProfile by the user ID.
	GetUserProfile(userID string, mediaType *app.UserProfile) error
	// UpdateUserProfile updates the UserProfile data for a particular user by its user ID.
	// If the profile already exists, it updates the data. If not profile entry exists, a new one is created.
	// Returns the updated or newly created user profile.
	UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error)
	UpdateMyProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error)
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

// GetUserProfile finds user profile by Id. Return media type if succeed.
func (c *MongoCollection) GetUserProfile(userID string, mediaType *app.UserProfile) error {
	objectUserID, err := hexToObjectID(userID)
	if err != nil {
		return err
	}

	if err := c.Find(bson.M{"userid": objectUserID}).One(&mediaType); err != nil {
		if err.Error() == "not found" {
			return goa.ErrNotFound(err)
		}
		return goa.ErrInternal(err)
	}

	return nil
}

// UpdateMyProfile updates own profile.
func (c *MongoCollection) UpdateMyProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {
	// Update(selector interface{}, update interface{}) error
	objectUserID, err := hexToObjectID(userID)
	if err != nil {
		return nil, err
	}

	err = c.Update(
		bson.M{"userid": objectUserID},
		bson.M{"$set": bson.M{
			"email":    profile.Email,
			"fullname": profile.FullName,
		},
		})

	// Handle errors
	if err != nil {
		if err.Error() == "not found" {
			return nil, goa.ErrNotFound(err)
		}
		return nil, goa.ErrInternal(err)
	}

	res := &app.UserProfile{
		UserID:   userID,
		FullName: &profile.FullName,
		Email:    &profile.Email,
	}

	return res, nil
}

// UpdateUserProfile updates user profile by id. Return media type if succeed.
func (c *MongoCollection) UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {
	objectUserID, err := hexToObjectID(userID)
	if err != nil {
		return nil, err
	}

	created := int(time.Now().Unix())
	_, err = c.Upsert(
		bson.M{"userid": objectUserID},
		bson.M{"$set": bson.M{
			"userid":    objectUserID,
			"email":     profile.Email,
			"fullname":  profile.FullName,
			"createdon": created,
		},
		})

	// Handle errors
	if err != nil {
		if mgo.IsDup(err) {
			return nil, goa.ErrBadRequest("email already exists in the database")
		}
		return nil, goa.ErrInternal(err)
	}

	res := &app.UserProfile{
		UserID:    userID,
		FullName:  &profile.FullName,
		Email:     &profile.Email,
		CreatedOn: created,
	}

	return res, nil
}

func hexToObjectID(hexID string) (bson.ObjectId, error) {
	// Return whether userID is a valid hex representation of object id.
	if bson.IsObjectIdHex(hexID) != true {
		return "", goa.ErrBadRequest("invalid User ID")
	}

	// Return an ObjectId from the provided hex representation.
	objectID := bson.ObjectIdHex(hexID)

	// Return true if objectID is valid. A valid objectID must contain exactly 12 bytes.
	if objectID.Valid() != true {
		return "", goa.ErrInternal("invalid object ID")
	}

	return objectID, nil
}
