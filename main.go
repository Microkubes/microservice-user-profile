//go:generate goagen bootstrap -d github.com/Microkubes/microservice-user-profile/design

package main

import (
	"fmt"
	"os"

	"github.com/Microkubes/microservice-security/chain"
	"github.com/Microkubes/microservice-security/flow"

	// "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/backends"
	toolscfg "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/microservice-tools/utils/healthcheck"
	"github.com/Microkubes/microservice-tools/utils/version"
	"github.com/Microkubes/microservice-user-profile/app"
	"github.com/Microkubes/microservice-user-profile/db"
	"github.com/keitaroinc/goa"
	"github.com/keitaroinc/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("user-profile")

	configFile := loadConfigSettings()

	cfg, err := toolscfg.LoadConfig(configFile)
	if err != nil {
		service.LogError("config", "err", err)
		return
	}

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

	service.Use(healthcheck.NewCheckMiddleware("/healthcheck"))

	service.Use(version.NewVersionMiddleware(cfg.Version, "/version"))

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
		"name": "user-profile",
		"indexes": []backends.Index{
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

func setupUserService(serviceConfig *toolscfg.ServiceConfig) (db.UserProfileRepository, error) {
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

func loadConfigSettings() string {
	serviceConfigFile := os.Getenv("SERVICE_CONFIG_FILE")

	if serviceConfigFile == "" {
		serviceConfigFile = "/run/secrets/microservice_user_profile_config.json"
	}

	return serviceConfigFile
}
