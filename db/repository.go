package db

import (
	"github.com/Microkubes/microservice-user-profile/app"
	// "github.com/goadesign/goa"

	// "gopkg.in/mgo.v2"
	// "time"
	backends "github.com/JormungandrK/backends"
	"fmt"
	
)

// User is an object which holds the UserID, FullName, Email and the date of creation 
type User struct {
	UserID 		string `json:"userId" bson:"userId"`
	FullName 	string `json:"fullname,omitempty" bson:"fullName,omitempty"`
	Email 		string `json:"email,omitempty" bson:"email,omitempty"`
	CreatedOn 	int    `json:"createdOn,omitempty" bson:"createdOn"` 
}


// UserProfileRepository defines the interface for accessing the user profile data
type UserProfileRepository interface {
	// GetUserProfile looks up a UserProfile by the user ID.
	GetUserProfile(userID string) (*app.UserProfile, error)
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
func (r *BackendUserService) GetUserProfile(userID string) (*app.UserProfile, error) {
	profile, err := r.userRepository.GetOne(backends.NewFilter().Match("userid", userID), &User{})
	if err != nil {
		return nil, err
	}
	result := &app.UserProfile{}
	if err = backends.MapToInterface(profile, result); 
	err != nil {
		return nil, err
	}
	return result, err
}


// UpdateUserProfile updates user profile by id. Return media type if succeed.
func (r *BackendUserService) UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {

	exitingIntf, err := r.userRepository.GetOne(backends.NewFilter().Match("userid", userID), &app.UserProfile{}) 
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		if !backends.IsErrNotFound(err){
			return nil, err			
		}
	}

	var existing *app.UserProfile

	var filter backends.Filter

	if exitingIntf == nil {
		existing = &app.UserProfile{
			FullName: &profile.FullName,
			Email: &profile.Email,
			UserID: userID,
		}
	} else {
		existing = exitingIntf.(*app.UserProfile)
		existing.FullName = &profile.FullName
		existing.Email = &profile.Email
		filter = backends.NewFilter().Match("userid", userID)
	}
	

	_, err = r.userRepository.Save(existing, filter)
	if err != nil {
		return nil, err
	}    
	    
	return existing, nil
}
