package controller

import "github.com/labstack/echo/v4"

type Topic interface {
	Post(echo.Context)
}
