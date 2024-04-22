package handlers

import (
	s "greet-sm-srv/services"

	m "github.com/greetinc/greet-middlewares/middlewares"

	"github.com/labstack/echo/v4"
)

type SocmedHandler interface {
	GetAll(c echo.Context) error
	// GetById(c echo.Context) error
	Create(c echo.Context) error
	// Update(c echo.Context) error
	// Delete(c echo.Context) error
}

type domainHandler struct {
	domainService s.SocmedService
	jwt           m.JWTService
}

func NewSocmedHandler(SocmedS s.SocmedService, jwtS m.JWTService) SocmedHandler {
	return &domainHandler{
		domainService: SocmedS,
		jwt:           jwtS,
	}
}
