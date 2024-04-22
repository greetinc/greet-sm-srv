package routes

import (
	"greet-sm-srv/configs"

	"github.com/greetinc/greet-middlewares/middlewares"
	"github.com/labstack/echo/v4"

	h_sm "greet-sm-srv/handlers"
	r_sm "greet-sm-srv/repositories"
	s_sm "greet-sm-srv/services"
)

var (
	DB = configs.InitDB()

	JWT   = middlewares.NewJWTService()
	blogR = r_sm.NewSocmedRepository(DB)
	blogS = s_sm.NewSocmedService(blogR, JWT)
	blogH = h_sm.NewSocmedHandler(blogS, JWT)
)

func New() *echo.Echo {

	e := echo.New()
	v1 := e.Group("api/v1")
	{
		sm := v1.Group("/blog", middlewares.AuthorizeJWT(JWT))
		{
			sm.POST("", blogH.Create)
			sm.GET("", blogH.GetAll)
		}
	}
	return e
}
