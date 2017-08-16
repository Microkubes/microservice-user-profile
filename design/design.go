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
	Host("localhost:8080")
})

var _ = Resource("userProfile", func() {
	BasePath("user-profile")
	DefaultMedia(UserProfileMedia)

	Action("GetUserProfile", func() {
		Description("Get a UserProfile by UserID")
		Routing(GET("/:userId"))
		Params(func() {
			Param("userId", String, "The user ID")
		})
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("GetMyProfile", func() {
		Description("Get a UserProfile by UserID")
		Routing(GET("/me"))
        Params(func() {
            Param("userId", String, "User ID")
            Required("userId")
        })		
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
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
	Attribute("userId", String, "Unique user id")
	Attribute("fullName", String, "Full name of the user")
	Attribute("email", String, "Email of user", func() {
		Format("email")
	})
})

// Swagger UI
var _ = Resource("swagger", func() {
	Description("The API swagger specification")

	Files("swagger.json", "swagger/swagger.json")
	Files("swagger-ui/*filepath", "swagger-ui/dist")
})
