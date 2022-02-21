package router

import (
	"github.com/labstack/echo"
	"xbrother/server/apis/board"
)

func RoutingDKMK(grp *echo.Group) {
	grp.GET("/board/list", board.ProcessBoardList)
}
