package main

import (
	"os"
	"os/signal"

	"github.com/aghaghiamh/ava/config"
	"github.com/aghaghiamh/ava/handler/httpserver"
	"github.com/aghaghiamh/ava/handler/httpserver/userhandler"

	storage "github.com/aghaghiamh/ava/repository"
	userdb "github.com/aghaghiamh/ava/repository/user"
	"github.com/aghaghiamh/ava/service/userservice"
	"github.com/aghaghiamh/ava/validator/uservalidator"
)

func main() {
	config := config.LoadConfig()

	// General DB Connector
	generalMysqlDB, _ := storage.New(config.DB)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// run the http server
	userHandler := setup(config, generalMysqlDB)
	server := httpserver.New(config.Server, userHandler)
	go func() {
		server.Serve()
	}()

	// Graceful Termination - wait until there is a os.signal on the quit channel then revoke all other children.
	<-quit

	server.Shutdown()
}

func setup(config config.Config, mysqlDB *storage.MysqlDB) userhandler.Handler {
	// User Service
	userRepo := userdb.New(mysqlDB)
	uservalidator := uservalidator.New(userRepo)
	userSvc := userservice.New(userRepo)
	userHandler := userhandler.New(config.HandlerConfig, uservalidator, userSvc)

	return userHandler
}
