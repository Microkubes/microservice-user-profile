package db

import (
	"github.com/Microkubes/microservice-user-profile/app"
	// "github.com/keitaroinc/goa"

	// "gopkg.in/mgo.v2"
	// "time"
	"fmt"

	backends "github.com/Microkubes/backends"
)

// Filter holds the exact Match values and Pattern match values for organizations lookup.
type Filter struct {
	// Match exact values.
	Match map[string]interface{}

	// Pattern match values by regular expression.
	Pattern map[string]string
}

// Sort holds the sorting specification for the results. This usually entails specifying
// the property by which to sort and the sorting direction (either "asc" or "desc").
type Sort struct {
	// SortBy is the name of the property by which to sort.
	SortBy string

	// Direction specifies the sorting direction. Valid values are: "asc" and "desc".
	Direction string
}

// UserProfilePage represents a paginated result of user profiles.
type UserProfilePage struct {
	// Page number. Starts from 1.
	Page int `json:"page"`

	// PageSize is the number of results per page.
	PageSize int `json:"pageSize"`

	// Items is an array of the actually returned user profiles that match the search filter.
	Items []User `json:"items"`
}

// User is an object which holds the UserID, FullName, Email and the date of creation
type User struct {
	UserID                    string `json:"userId" bson:"userId"`
	FullName                  string `json:"fullname,omitempty" bson:"fullName,omitempty"`
	Email                     string `json:"email,omitempty" bson:"email,omitempty"`
	Company                   string `json:"company,omitempty" bson:"company,omitempty"`
	CompanyRegistrationNumber string `json:"companyRegistrationNumber,omitempty" bson:"companyRegistrationNumber,omitempty"`
	TaxNumber                 string `json:"taxNumber,omitempty" bson:"taxNumber,omitempty"`
	CreatedOn                 int    `json:"createdOn,omitempty" bson:"createdOn"`
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
	// Find Performs a search for user profiles that match the input filter.
	Find(filter *Filter, sort *Sort, page, pageSize int) (*UserProfilePage, error)
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
	profile, err := r.userRepository.GetOne(backends.NewFilter().Match("userId", userID), &User{})
	if err != nil {
		return nil, err
	}
	result := &app.UserProfile{}
	if err = backends.MapToInterface(profile, result); err != nil {
		return nil, err
	}
	return result, err
}

// UpdateUserProfile updates user profile by id. Return media type if succeed.
func (r *BackendUserService) UpdateUserProfile(profile *app.UserProfilePayload, userID string) (*app.UserProfile, error) {

	exitingIntf, err := r.userRepository.GetOne(backends.NewFilter().Match("userId", userID), &app.UserProfile{})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		if !backends.IsErrNotFound(err) {
			return nil, err
		}
	}

	var existing *app.UserProfile

	var filter backends.Filter

	if exitingIntf == nil {
		existing = &app.UserProfile{
			FullName:                  &profile.FullName,
			Email:                     &profile.Email,
			UserID:                    userID,
			Company:                   profile.Company,
			CompanyRegistrationNumber: profile.CompanyRegistrationNumber,
			TaxNumber:                 profile.TaxNumber,
		}
	} else {
		existing = exitingIntf.(*app.UserProfile)
		existing.FullName = &profile.FullName
		existing.Email = &profile.Email
		existing.Company = profile.Company
		existing.CompanyRegistrationNumber = profile.CompanyRegistrationNumber
		existing.TaxNumber = profile.TaxNumber
		filter = backends.NewFilter().Match("userId", userID)
	}

	_, err = r.userRepository.Save(existing, filter)

	if err != nil {
		return nil, err
	}
	return existing, nil
}

// Find looks up the user profiles matching the search filter.
func (r *BackendUserService) Find(filter *Filter, sort *Sort, page, pageSize int) (*UserProfilePage, error) {
	var bf backends.Filter
	if filter != nil {
		bf = backends.NewFilter()
		if filter.Match != nil {
			for prop, value := range filter.Match {
				bf.Match(prop, value)
			}
		}
		if filter.Pattern != nil {
			for prop, pattern := range filter.Pattern {
				bf.MatchPattern(prop, pattern)
			}
		}
	}

	sortBy := "createdAt"
	sortDir := "asc"

	if sort != nil {
		sortBy = sort.SortBy
		sortDir = sort.Direction
	}

	res, err := r.userRepository.GetAll(bf, &User{}, sortBy, sortDir, pageSize, page)
	if err != nil {
		return nil, err
	}

	userPrfls := *(res.(*[]*User))
	userProfiles := []User{}
	for _, org := range userPrfls {
		userProfiles = append(userProfiles, *org)
	}

	return &UserProfilePage{
		Page:     page,
		PageSize: pageSize,
		Items:    userProfiles,
	}, nil
}
