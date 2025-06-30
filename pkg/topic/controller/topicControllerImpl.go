package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yuri7772d/reditCloneApi/pkg/topic/service"
)

type topicImpl struct {
	service service.Topic
}

func New(service service.Topic) Topic {
	return &topicImpl{service}
}
func (c *topicImpl) Post(pctx echo.Context) {

}
