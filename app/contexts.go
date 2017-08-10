// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "user-profile": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-user-profile/design
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-user-profile
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// GetUserProfileUserProfileContext provides the userProfile GetUserProfile action context.
type GetUserProfileUserProfileContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	UserID string
}

// NewGetUserProfileUserProfileContext parses the incoming request URL and body, performs validations and creates the
// context used by the userProfile controller GetUserProfile action.
func NewGetUserProfileUserProfileContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetUserProfileUserProfileContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetUserProfileUserProfileContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUserID := req.Params["userId"]
	if len(paramUserID) > 0 {
		rawUserID := paramUserID[0]
		rctx.UserID = rawUserID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetUserProfileUserProfileContext) OK(r *UserProfile) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/jormungandr.user-profile+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetUserProfileUserProfileContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}
