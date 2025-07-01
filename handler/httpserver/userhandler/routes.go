package userhandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/user")

	userGroup.POST("/register", h.RegisterHandler)
	userGroup.GET("/profile/:id", h.GetProfileHandler)
	userGroup.GET("/delete/:id", h.DeleteHandler)
}
