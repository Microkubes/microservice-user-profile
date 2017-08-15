//go:generate goagen bootstrap -d github.com/JormungandrK/microservice-user-profile/design

package main

import (
	"net/http"
	"os"

	"github.com/JormungandrK/microservice-tools/gateway"
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"gopkg.in/mgo.v2"
)

const (
 	Host     = "127.0.0.1:27017"
 	Username = "restapi"
	Password = "restapi"
	Database = "user-profile"
)

func main() {
	// Gateway self-registration
	unregisterService := registerMicroservice()
	defer unregisterService() // defer the unregister for after main exits

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	service := goa.New("user-profile")

	// Load MongoDB ENV variables
	host, username, password, database := loadMongnoSettings()
	// Create new session to MongoDB
	session := store.NewSession(host, username, password, database)

	// At the end close session
	defer session.Close()

	// Create users collection and indexes
	indexes := []string{"fullname", "email"}
	userProfileCollection := store.PrepareDB(session, database, "user-profiles", indexes)

	// Mount "swagger" controller
	c := NewSwaggerController(service)
	app.MountSwaggerController(service, c)
	// Mount "userProfile" controller
	c2 := NewUserProfileController(service, &db.UserProfileRepository{Collection: userProfileCollection})
	app.MountUserProfileController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func loadMongnoSettings() (string, string, string, string) {
	host := os.Getenv("MONGO_URL")
	username := os.Getenv("MS_USERNAME")
	password := os.Getenv("MS_PASSWORD")
	database := os.Getenv("MS_DBNAME")

	if host == "" {
		host = "127.0.0.1:27017"
	}
	if username == "" {
		username = "restapi"
	}
	if password == "" {
		password = "restapi"
	}
	if database == "" {
		database = "user-profile"
	}

	return host, username, password, database
}

func loadGatewaySettings() (string, string) {
	gatewayURL := os.Getenv("API_GATEWAY_URL")
	serviceConfigFile := os.Getenv("SERVICE_CONFIG_FILE")

	if gatewayURL == "" {
		gatewayURL = "http://localhost:8001"
	}
	if serviceConfigFile == "" {
		serviceConfigFile = "config.json"
	}

	return gatewayURL, serviceConfigFile
}

func registerMicroservice() func() {
	gatewayURL, configFile := loadGatewaySettings()
	registration, err := gateway.NewKongGatewayFromConfigFile(gatewayURL, &http.Client{}, configFile)
	if err != nil {
		panic(err)
	}
	err = registration.SelfRegister()
	if err != nil {
		panic(err)
	}

	return func() {
		registration.Unregister()
	}
}
