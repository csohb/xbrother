package context

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"xbrother/common/api_context"
	"xbrother/server/config"
)

type Context struct {
	*api_context.Context
	Config     *config.DKMKConfig
	DB         *gorm.DB
	httpClient *resty.Client
}

func InitContext(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, client *resty.Client, cfg *config.DKMKConfig) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := &Context{
				Context:    api_context.NewContext(context, logger),
				Config:     cfg,
				DB:         db,
				httpClient: client,
			}
			return next(ctx)
		}
	})
}
