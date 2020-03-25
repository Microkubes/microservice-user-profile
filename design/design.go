package design

import (
	. "github.com/keitaroinc/goa/design"
	. "github.com/keitaroinc/goa/design/apidsl"
)

var _ = API("user-profile", func() {
	Title("User Profile Microservice")
	Description("API for managing UserProfile data.")
	Version("1.0")
	Scheme("http")
	Host("localhost:8082")
})

var _ = Resource("userProfile", func() {
	BasePath("profiles")
	DefaultMedia(UserProfileMedia)

	// Allow preflight HTTP requests
	Origin("*", func() {
		Methods("OPTIONS")
	})

	Action("GetUserProfile", func() {
		Description("Get a UserProfile by UserID")
		Routing(GET(":userId"))
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
		Routing(PUT("me"))
		Payload(UserProfilePayload)
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("GetMyProfile", func() {
		Description("Get a UserProfile by UserID")
		Routing(GET("me"))
		Response(OK)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("UpdateUserProfile", func() {
		Description("Update user profile")
		Routing(PUT(":userId"))
		Params(func() {
			Param("userId", String, "User ID")
		})
		Payload(UserProfilePayload)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
		Response(OK, UserProfileMedia)
	})

	Action("FindUserProfile", func() {
		Description("Find (filter) organizations by some filter.")
		Routing(POST("/find"))
		Payload(FilterPayload)
		Response(OK, UserProfilePageMedia)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})
})

// UserProfileMedia is the default media type for user-profile service
var UserProfileMedia = MediaType("application/microkubes.user-profile+json", func() {
	TypeName("userProfile")
	Reference(UserProfilePayload)

	Attributes(func() {
		Attribute("userId", String, "Unique user ID")
		Attribute("fullName")
		Attribute("email")
		Attribute("company")
		Attribute("companyRegistrationNumber")
		Attribute("taxNumber")
		Attribute("createdOn", Integer, "User profile created timestamp")
		Required("userId", "createdOn")
	})

	View("default", func() {
		Attribute("userId")
		Attribute("fullName")
		Attribute("email")
		Attribute("company")
		Attribute("companyRegistrationNumber")
		Attribute("taxNumber")
		Attribute("createdOn")
	})
})

// UserProfilePageMedia result of filter-by. One result page along with items (array of UserProfiles).
var UserProfilePageMedia = MediaType("application/microkubes.user-profile-page+json", func() {
	TypeName("UserProfilePage")
	Attributes(func() {
		Attribute("page", Integer, "Page number (1-based).")
		Attribute("pageSize", Integer, "Items per page.")
		Attribute("items", ArrayOf(UserProfileMedia), "User profile list")
	})
	View("default", func() {
		Attribute("page")
		Attribute("pageSize")
		Attribute("items")
	})
})

// UserProfilePayload is the payload specification
var UserProfilePayload = Type("UserProfilePayload", func() {
	Description("UserProfile data")

	Attribute("fullName", String, "Full name of the user")
	Attribute("email", String, "Email of user", func() {
		Format("email")
	})
	Attribute("company", String, "Company name")
	Attribute("companyRegistrationNumber", String, "Company registration number")
	Attribute("taxNumber", String, "Tax number")

	Required("fullName", "email")
})

// FilterPayload Organizations filter request payload.
var FilterPayload = Type("FilterPayload", func() {
	Attribute("page", Integer, "Page number (1-based).")
	Attribute("pageSize", Integer, "Items per page.")
	Attribute("filter", ArrayOf(FilterProperty), "Organizations filter.")
	Attribute("sort", OrderSpec, "Sort specification.")
	Required("page", "pageSize")
})

// FilterProperty Single property filter. Holds the property name and the value to be matched for that property.
var FilterProperty = Type("FilterProperty", func() {
	Attribute("property", String, "Property name")
	Attribute("value", String, "Property value to match")
	Required("property", "value")
})

// OrderSpec specifies the sorting - by which property and the direction, either 'asc' (ascending)
// or 'desc' (descending).
var OrderSpec = Type("OrderSpec", func() {
	Attribute("property", String, "Sort by property")
	Attribute("direction", String, "Sort order. Can be 'asc' or 'desc'.")
	Required("property", "direction")
})

// Swagger UI
var _ = Resource("swagger", func() {
	Description("The API swagger specification")

	Files("swagger.json", "swagger/swagger.json")
	Files("swagger-ui/*filepath", "swagger-ui/dist")
})
