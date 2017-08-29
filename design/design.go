package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("user-profile", func() {
	Title("User Profile Microservice")
	Description("API for managing UserProfile data.")
	Version("1.0")
	Scheme("http")
	Host("localhost:8082")
})

var _ = Resource("userProfile", func() {
	DefaultMedia(UserProfileMedia)

	Action("GetUserProfile", func() {
		Description("Get a UserProfile by UserID")
		Routing(GET("users/:userId/profile"))
		Params(func() {
			Param("userId", String, "The user ID")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("UpdateMyProfile", func() {
		Description("Update my profile")
		Routing(PUT("profiles/me"))
		Payload(UserProfilePayload)
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("GetMyProfile", func() {
		Description("Get a UserProfile by UserID")
		Routing(GET("profiles/me"))
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("UpdateUserProfile", func() {
		Description("Update user profile")
		Routing(PUT("users/:userId/profile"))
		Params(func() {
			Param("userId", String, "User ID")
		})
		Payload(UserProfilePayload)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
		Response(OK, UserProfileMedia)
	})
})

// UserProfileMedia is the default media type for user-profile service
var UserProfileMedia = MediaType("application/jormungandr.user-profile+json", func() {
	TypeName("userProfile")
	Reference(UserProfilePayload)

	Attributes(func() {
		Attribute("userId", String, "Unique user ID")
		Attribute("fullName")
		Attribute("email")
		Attribute("createdOn", Integer, "User profile created timestamp")
		Required("userId", "createdOn")
	})

	View("default", func() {
		Attribute("userId")
		Attribute("fullName")
		Attribute("email")
		Attribute("createdOn")
	})
})

// UserProfilePayload is the payload specification
var UserProfilePayload = Type("UserProfilePayload", func() {
	Description("UserProfile data")

	Attribute("fullName", String, "Full name of the user")
	Attribute("email", String, "Email of user", func() {
		Format("email")
	})

	Required("fullName", "email")
})

// Swagger UI
var _ = Resource("swagger", func() {
	Description("The API swagger specification")

	Files("swagger.json", "swagger/swagger.json")
	Files("swagger-ui/*filepath", "swagger-ui/dist")
})
