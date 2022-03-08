package route

import "github.com/labstack/echo/v4"

func play(e *echo.Echo) {
	plays := e.Group("/plays")
	plays.GET("/with-bot", nil)
	plays.GET("/history", nil)
}
