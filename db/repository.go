package db

import (
	"github.com/Microkubes/microservice-user-profile/app"
	// "github.com/goadesign/goa"

	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "time"
	"github.com/JormungandrK/backends"
)

// User is an object which holds the UserID, FullName, Email and the date of creation 
type User struct {
	UserID 		string `json:"userid" bson:"userid"`
	FullName 	string `json:"fullname" bson:"fullname"`
	Email 		string `json:"email" bson:"email"`
	CreatedOn 	int    `json:"createdOn,omitempty" bson:"createdOn"` 
}


// UserProfileRepository defines the interface for accessing the user profile data
type UserProfileRepository interface {
	// GetUserProfile looks up a UserProfile by the user ID.
	GetUserProfile(userID string, mediaType *app.UserProfile) error
	// UpdateUserProfile updates the UserProfile data for a particular user by its user ID.
	// If the profile already exists, it updates the data. If not profile entry exists, a new one is created.
	// Returns the updated or newly created user profile.
	UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error)
	// UpdateMyProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error)
}



type BackendUserService struct {
	userRepository backends.Repository
}

func NewUserService(userRepository backends.Repository) UserProfileRepository {
	return &BackendUserService{
		userRepository: userRepository,
	}
}
 
// GetUserProfile finds user profile by Id. Return media type if succeed.
func (r *BackendUserService) GetUserProfile(userID string, mediaType *app.UserProfile) error {
	_, err := r.userRepository.GetOne(backends.NewFilter().Match("id", userID), mediaType)
	if err != nil {
		return err
	}
	// return result.(*app.UserProfile), nil
	return err
}

// UpdateUserProfile updates user profile by id. Return media type if succeed.
func (r *BackendUserService) UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {

	_, err := r.userRepository.GetOne(backends.NewFilter().Match("id", userID), &app.UserProfile{}) 
	if err != nil {
		return nil, err
	}

	existing := &app.UserProfile{}

	_, err = r.userRepository.Save(existing, backends.NewFilter().Match("id", userID))
	if err != nil {
		return nil, err
	}    
	    
	return existing, nil
}
