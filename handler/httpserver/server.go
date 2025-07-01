package httpserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aghaghiamh/ava/handler/httpserver/userhandler"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpConfig struct {
	Host                    string        `mapstructure:"host"`
	Port                    string        `mapstructure:"port"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type Server struct {
	cfg               HttpConfig
	router            *echo.Echo
	userHandler       userhandler.Handler
}

func New(cfg HttpConfig, userHandler userhandler.Handler) Server {
	return Server{
		cfg:               cfg,
		router:            echo.New(),
		userHandler:       userHandler,
	}
}

func (s *Server) Serve() {
	s.router.Use(middleware.Recover())

	s.userHandler.SetRoutes(s.router)

	if err := s.router.Start(s.cfg.Host + ":" + s.cfg.Port); err != nil {
		log.Fatalf("Couldn't Listen to the %s port.", s.cfg.Port)
	}
}

func (s Server) Shutdown() {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("error while shutting down the server: ", err)
	}

	fmt.Println("Gracefully shutdowned!!")
	<-ctxWithTimeout.Done()
}
