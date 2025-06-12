package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/yuri7772d/reditCloneApi/config"
	"gorm.io/gorm"
)

type echoserver struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

var (
	once   *sync.Once
	server *echoserver
)

func (s *echoserver) Start() {
	s.app.GET("/v1/health", s.healthCheck)
	s.httpListing()

}

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoserver {

	once = &sync.Once{}
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

func (s *echoserver) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
