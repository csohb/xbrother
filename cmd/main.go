package main

import (
	"flag"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"xbrother/databases/mysql"
	"xbrother/server/apis/context"
	"xbrother/server/config"
	"xbrother/server/router"
)

var (
	configFile = flag.String("c", "dkmk.yaml", "The path to the config file.")
)

func main() {
	flag.Parse()
	cfg := config.LoadDKMKConfig(*configFile)

	e := echo.New()
	defer e.Close()

	db, err := mysql.Connect(cfg.DB.DSN, cfg.DB.MaxIdle, cfg.DB.MaxConn)
	if err != nil {
		logrus.WithError(err).Error("db connect failed.")
		os.Exit(-1)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          nil,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "OPTION", "PUT", "HEAD"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    nil,
		MaxAge:           0,
	}))

	grp := e.Group(".api/v1")
	request := resty.New()
	context.InitContext(e, logrus.StandardLogger(), db, request, cfg)

	router.RoutingDKMK(grp)

	if err := e.Start(cfg.Server.Listen); err != nil {
		logrus.Panic("server start failed : ", err)
	}

	logrus.Exit(0)
}
