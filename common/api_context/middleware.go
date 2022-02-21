package api_context

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"time"
)

func NewContext(c echo.Context, logger *logrus.Logger) *Context {
	c.Set("requestTime", time.Now())
	return &Context{
		Context:     c,
		log:         logger,
		Log:         logger.WithField("uri", c.Request().RequestURI),
		requestTime: time.Time{},
	}
}
