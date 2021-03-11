package router

import (
	"net/http"

	"github.com/Topi99/academy-go-q12021/interface/controller"
	"github.com/labstack/echo"
)

// HTTPErrorHandler customize echo's HTTP error handler.
func HTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		key  = "ServerError"
		msg  string
	)

	if he, ok := err.(*controller.HTTPError); ok {
		code = he.Code
		key = he.Key
		msg = he.Message
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, controller.NewHTTPError(code, key, msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
