package echocors

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/rs/cors"
)

func Init(e *echo.Echo, Debug bool) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowCredentials: true,
		Debug:            Debug,
	})
	e.Use(standard.WrapMiddleware(c.Handler))
}
