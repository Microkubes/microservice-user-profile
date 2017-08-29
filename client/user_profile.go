// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "user-profile": userProfile Resource Client
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-user-profile/design
// --out=$(GOPATH)src/github.com/JormungandrK/microservice-user-profile
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// GetMyProfileUserProfilePath computes a request path to the GetMyProfile action of userProfile.
func GetMyProfileUserProfilePath() string {

	return fmt.Sprintf("/profiles/me")
}

// Get a UserProfile by UserID
func (c *Client) GetMyProfileUserProfile(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetMyProfileUserProfileRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetMyProfileUserProfileRequest create the request corresponding to the GetMyProfile action endpoint of the userProfile resource.
func (c *Client) NewGetMyProfileUserProfileRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetUserProfileUserProfilePath computes a request path to the GetUserProfile action of userProfile.
func GetUserProfileUserProfilePath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/users/%s/profile", param0)
}

// Get a UserProfile by UserID
func (c *Client) GetUserProfileUserProfile(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserProfileUserProfileRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserProfileUserProfileRequest create the request corresponding to the GetUserProfile action endpoint of the userProfile resource.
func (c *Client) NewGetUserProfileUserProfileRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateMyProfileUserProfilePath computes a request path to the UpdateMyProfile action of userProfile.
func UpdateMyProfileUserProfilePath() string {

	return fmt.Sprintf("/profiles/me")
}

// Update my profile
func (c *Client) UpdateMyProfileUserProfile(ctx context.Context, path string, payload *UserProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateMyProfileUserProfileRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateMyProfileUserProfileRequest create the request corresponding to the UpdateMyProfile action endpoint of the userProfile resource.
func (c *Client) NewUpdateMyProfileUserProfileRequest(ctx context.Context, path string, payload *UserProfilePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// UpdateUserProfileUserProfilePath computes a request path to the UpdateUserProfile action of userProfile.
func UpdateUserProfileUserProfilePath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/users/%s/profile", param0)
}

// Update user profile
func (c *Client) UpdateUserProfileUserProfile(ctx context.Context, path string, payload *UserProfilePayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUserProfileUserProfileRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserProfileUserProfileRequest create the request corresponding to the UpdateUserProfile action endpoint of the userProfile resource.
func (c *Client) NewUpdateUserProfileUserProfileRequest(ctx context.Context, path string, payload *UserProfilePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
