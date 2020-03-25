// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "user-profile": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-user-profile/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-user-profile
// --version=v1.3.1

package app

import (
	"context"
	"net/http"

	"github.com/keitaroinc/goa"
	"github.com/keitaroinc/goa/cors"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/dist")
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/dist/index.html")
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist/index.html", "route", "GET /swagger-ui/")
}

// UserProfileController is the controller interface for the UserProfile actions.
type UserProfileController interface {
	goa.Muxer
	FindUserProfile(*FindUserProfileUserProfileContext) error
	GetMyProfile(*GetMyProfileUserProfileContext) error
	GetUserProfile(*GetUserProfileUserProfileContext) error
	UpdateMyProfile(*UpdateMyProfileUserProfileContext) error
	UpdateUserProfile(*UpdateUserProfileUserProfileContext) error
}

// MountUserProfileController "mounts" a UserProfile resource controller on the given service.
func MountUserProfileController(service *goa.Service, ctrl UserProfileController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/profiles/find", ctrl.MuxHandler("preflight", handleUserProfileOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/profiles/me", ctrl.MuxHandler("preflight", handleUserProfileOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/profiles/:userId", ctrl.MuxHandler("preflight", handleUserProfileOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewFindUserProfileUserProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*FilterPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.FindUserProfile(rctx)
	}
	h = handleUserProfileOrigin(h)
	service.Mux.Handle("POST", "/profiles/find", ctrl.MuxHandler("FindUserProfile", h, unmarshalFindUserProfileUserProfilePayload))
	service.LogInfo("mount", "ctrl", "UserProfile", "action", "FindUserProfile", "route", "POST /profiles/find")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetMyProfileUserProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetMyProfile(rctx)
	}
	h = handleUserProfileOrigin(h)
	service.Mux.Handle("GET", "/profiles/me", ctrl.MuxHandler("GetMyProfile", h, nil))
	service.LogInfo("mount", "ctrl", "UserProfile", "action", "GetMyProfile", "route", "GET /profiles/me")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserProfileUserProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetUserProfile(rctx)
	}
	h = handleUserProfileOrigin(h)
	service.Mux.Handle("GET", "/profiles/:userId", ctrl.MuxHandler("GetUserProfile", h, nil))
	service.LogInfo("mount", "ctrl", "UserProfile", "action", "GetUserProfile", "route", "GET /profiles/:userId")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateMyProfileUserProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserProfilePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateMyProfile(rctx)
	}
	h = handleUserProfileOrigin(h)
	service.Mux.Handle("PUT", "/profiles/me", ctrl.MuxHandler("UpdateMyProfile", h, unmarshalUpdateMyProfileUserProfilePayload))
	service.LogInfo("mount", "ctrl", "UserProfile", "action", "UpdateMyProfile", "route", "PUT /profiles/me")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUserProfileUserProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserProfilePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateUserProfile(rctx)
	}
	h = handleUserProfileOrigin(h)
	service.Mux.Handle("PUT", "/profiles/:userId", ctrl.MuxHandler("UpdateUserProfile", h, unmarshalUpdateUserProfileUserProfilePayload))
	service.LogInfo("mount", "ctrl", "UserProfile", "action", "UpdateUserProfile", "route", "PUT /profiles/:userId")
}

// handleUserProfileOrigin applies the CORS response headers corresponding to the origin.
func handleUserProfileOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalFindUserProfileUserProfilePayload unmarshals the request body into the context request data Payload field.
func unmarshalFindUserProfileUserProfilePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &filterPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateMyProfileUserProfilePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateMyProfileUserProfilePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userProfilePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserProfileUserProfilePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserProfileUserProfilePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userProfilePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
