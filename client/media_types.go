// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user-profile": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-user-profile/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-user-profile
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// userProfile media type (default view)
//
// Identifier: application/microkubes.user-profile+json; view=default
type UserProfile struct {
	// User profile created timestamp
	CreatedOn int `form:"createdOn" json:"createdOn" yaml:"createdOn" xml:"createdOn"`
	// Email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty" xml:"email,omitempty"`
	// Full name of the user
	FullName *string `form:"fullName,omitempty" json:"fullName,omitempty" yaml:"fullName,omitempty" xml:"fullName,omitempty"`
	// Unique user ID
	UserID string `form:"userId" json:"userId" yaml:"userId" xml:"userId"`
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

// DecodeUserProfile decodes the UserProfile instance encoded in resp body.
func (c *Client) DecodeUserProfile(resp *http.Response) (*UserProfile, error) {
	var decoded UserProfile
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
