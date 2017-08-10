// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "user-profile": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-user-profile/design
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-user-profile
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
)

// userProfile media type (default view)
//
// Identifier: application/jormungandr.user-profile+json; view=default
type UserProfile struct {
	// User profile created timestamp
	CreatedOn int `form:"createdOn" json:"createdOn" xml:"createdOn"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Full name of the user
	FullName *string `form:"fullName,omitempty" json:"fullName,omitempty" xml:"fullName,omitempty"`
	// Unique user ID
	UserID string `form:"userId" json:"userId" xml:"userId"`
}

// Validate validates the UserProfile media type instance.
func (mt *UserProfile) Validate() (err error) {
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}

	if mt.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *mt.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, *mt.Email, goa.FormatEmail, err2))
		}
	}
	return
}
