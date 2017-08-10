package db

import (
	"github.com/JormungandrK/microservice-user-profile/app"
)

// UserProfileRepository defaines the interface for accessing the user profile data
type UserProfileRepository interface {
	// GetUserProfile looks up a UserProfile by the user ID. Returns app.UserProfile if found or nil if not found.
	GetUserProfile(userID string) (*app.UserProfile, error)

	// UpdateUserProfile updates the UserProfile data for a particular user by its user ID.
	// If the profile already exists, it updates the data. If not profile entry exists, a new one is created.
	// Returns the updated or newly created user profile.
	UpdateUserProfile(profile app.UserProfilePayload) (*app.UserProfile, error)
}
