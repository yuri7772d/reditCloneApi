package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/yuri7772d/reditCloneApi/config"
	"github.com/yuri7772d/reditCloneApi/databases"
)

type echoserver struct {
	app  *echo.Echo
	db   databases.Database
	conf *config.Config
}

var (
	once   sync.Once
	server *echoserver
)

func (s *echoserver) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(getBodyLimit(s.conf.Server.BodyLimit))
	s.app.Use(getCORSMiddleware(s.conf.Server.AllowOrigins))
	s.app.Use(getTimeoutMiddleware(s.conf.Server.Timeout))

	s.app.GET("/v1/health", s.healthCheck)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefullyshutdown(quitCh)

	s.httpListing()

}

func NewEchoServer(conf *config.Config, db databases.Database) *echoserver {

	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	once.Do(func() {
		server = &echoserver{
			app:  echoApp,
			db:   db,
			conf: conf,
		}
	})
	return server
}
func (s *echoserver) httpListing() {
	url := fmt.Sprintf(":%d", s.conf.Server.Port)
	if err := s.app.Start(url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatalf("Error : %s", err.Error())
	}
}
func (s *echoserver) gracefullyshutdown(quitCh chan os.Signal) {
	<-quitCh
	s.app.Logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatalf("Error : %s", err.Error())
	}
}

func getTimeoutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Request timeout",
		Timeout:      timeout * time.Second,
	})
}
func getCORSMiddleware(allawOrigings []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allawOrigings,
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.DELETE, echo.POST},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderAccept, echo.HeaderContentType},
	})
}
func getBodyLimit(bobyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bobyLimit)
}

func (s *echoserver) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
