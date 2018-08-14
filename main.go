//go:generate goagen bootstrap -d github.com/Microkubes/microservice-user-profile/design

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Microkubes/microservice-security/chain"
	"github.com/Microkubes/microservice-security/flow"
	// "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/microservice-tools/gateway"
	"github.com/Microkubes/microservice-user-profile/app"
	"github.com/Microkubes/microservice-user-profile/db"
	toolscfg "github.com/Microkubes/microservice-tools/config"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/JormungandrK/backends"
)

func main() {
	// Create service
	service := goa.New("user-profile")

	gatewayAdminURL, configFile := loadGatewaySettings()

	cfg, err := toolscfg.LoadConfig(configFile)
	if err != nil {
		service.LogError("config", "err", err)
	}
	// Gateway self-registration
	unregisterService := registerMicroservice(gatewayAdminURL, cfg)
	defer unregisterService() // defer the unregister for after main exits

	// Setup user-profile service
	userService, err := setupUserService(cfg)
	if err != nil {
		service.LogError("config", err)
		return
	}

	securityChain, cleanup, err := flow.NewSecurityFromConfig(cfg)
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

	// Mount "user-profile" controller
	c := NewUserProfileController(service, userService)
	app.MountUserProfileController(service, c)

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%d", cfg.Service.MicroservicePort)); err != nil {
		service.LogError("startup", "err", err)
	}

}

func setupRepository(backend backends.Backend) (backends.Repository, error) {
	return backend.DefineRepository("user-profile", backends.RepositoryDefinitionMap{
		"name":    "user-profile",
		"indexes":  []backends.Index{
				backends.NewUniqueIndex("userId"),
				backends.NewUniqueIndex("fullname"),
		},
		"hashKey":       "id",
		"readCapacity":  int64(5),
		"writeCapacity": int64(5),
		"GSI": map[string]interface{}{
			"email": map[string]interface{}{
				"readCapacity":  1,
				"writeCapacity": 1,
			},
		},
	})
}


func setupBackend(dbConfig toolscfg.DBConfig) (backends.Backend, backends.BackendManager, error) {
	dbinfo := map[string]*toolscfg.DBInfo{}

	dbinfo[dbConfig.DBName] = &dbConfig.DBInfo

	backendsManager := backends.NewBackendSupport(dbinfo)
	backend, err := backendsManager.GetBackend(dbConfig.DBName)

	return backend, backendsManager, err 
}


func setupUserService(serviceConfig *toolscfg.ServiceConfig) (db.UserProfileRepository, error){
	backend, _, err := setupBackend(serviceConfig.DBConfig)
	if err != nil {
		return nil, err
	}

	userRepo, err := setupRepository(backend)
	if err != nil {
		return nil, err
	}

	return db.NewUserService(userRepo), err
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

func registerMicroservice(gatewayAdminURL string, cfg *toolscfg.ServiceConfig) func() {
	registration := gateway.NewKongGateway(gatewayAdminURL, &http.Client{}, cfg.Service)
	err := registration.SelfRegister()
	if err != nil {
		panic(err)
	}

	return func() {
		registration.Unregister()
	}
}
 