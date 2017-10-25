//go:generate goagen bootstrap -d github.com/JormungandrK/microservice-user-profile/design

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JormungandrK/microservice-security/chain"
	"github.com/JormungandrK/microservice-security/flow"
	"github.com/JormungandrK/microservice-tools/config"
	"github.com/JormungandrK/microservice-tools/gateway"
	"github.com/JormungandrK/microservice-user-profile/app"
	"github.com/JormungandrK/microservice-user-profile/db"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("user-profile")

	gatewayAdminURL, configFile := loadGatewaySettings()

	conf, err := config.LoadConfig(configFile)
	if err != nil {
		service.LogError("config", "err", err)
		return
	}
	// Gateway self-registration
	unregisterService := registerMicroservice(gatewayAdminURL, conf)
	defer unregisterService() // defer the unregister for after main exits

	securityChain, cleanup, err := flow.NewSecurityFromConfig(conf)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	service.Use(chain.AsGoaMiddleware(securityChain))

	// Load MongoDB ENV variables
	//host, username, password, database := loadMongnoSettings()
	dbConf := conf.DBConfig
	// Create new session to MongoDB
	session := db.NewSession(dbConf.Host, dbConf.Username, dbConf.Password, dbConf.DatabaseName)

	// At the end close session
	defer session.Close()

	// Create users collection and indexes
	indexes := []string{"email"}
	userProfileCollection := db.PrepareDB(session, dbConf.DatabaseName, "user-profiles", indexes)

	// Mount "swagger" controller
	c := NewSwaggerController(service)
	app.MountSwaggerController(service, c)
	// Mount "userProfile" controller
	c2 := NewUserProfileController(service, &db.MongoCollection{Collection: userProfileCollection})
	app.MountUserProfileController(service, c2)

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%d", conf.Service.MicroservicePort)); err != nil {
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
		database = "user-profiles"
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
		serviceConfigFile = "/run/secrets/microservice_user_profile_config.json"
	}

	return gatewayURL, serviceConfigFile
}

func registerMicroservice(gatewayAdminURL string, conf *config.ServiceConfig) func() {
	registration := gateway.NewKongGateway(gatewayAdminURL, &http.Client{}, conf.Service)
	err := registration.SelfRegister()
	if err != nil {
		panic(err)
	}

	return func() {
		registration.Unregister()
	}
}
