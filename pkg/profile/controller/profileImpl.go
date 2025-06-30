package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yuri7772d/reditCloneApi/pkg/topic/service"
)

type profileImpl struct {
	service service.Topic
}

func New(service service.Topic) Profile {
	return &profileImpl{}
}

func (p *profileImpl) Create(pctx echo.Context) {

}
