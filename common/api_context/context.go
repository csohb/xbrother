package api_context

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type ServiceHandler interface {
	Service() *CommonResponse
	GetRequestData() interface{}
	ApiName() interface{}
	GetContext() *Context
}

type Context struct {
	echo.Context
	log         *logrus.Logger
	Log         *logrus.Entry
	requestTime time.Time
	RespData    CommonResponse
}

func (ctx *Context) GetContext() *Context {
	return ctx
}

func (ctx *Context) SetLogAPI(name string) {
	ctx.Log = ctx.Log.WithField("api", name)
}

func Process(service ServiceHandler) error {
	var err error
	c := service.GetContext()
	c.Log = c.Log.WithField("api", service.ApiName())

	if c.Request().Method == http.MethodOptions {
		return c.NoContent(http.StatusOK)
	}

	reqData := service.GetRequestData()
	if reqData != nil {
		if err = c.Bind(reqData); err != nil {
			c.Log.Errorf("request decoding failed : %+v", err)
			return BadJSON("The request body could not be decoded into vaild JSON " + err.Error()).Send(c)
		}
	}

	resp := service.Service()
	if resp != nil {
		return resp.Send(c)
	}
	return nil
}
