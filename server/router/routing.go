package router

import (
	"github.com/labstack/echo"
	"xbrother/server/apis/board"
	"xbrother/server/apis/lines"
	"xbrother/server/apis/vote"
)

func RoutingDKMK(grp *echo.Group) {
	grp.GET("/board/list", board.ProcessBoardList)
	grp.POST("/board/insert", board.ProcessBoardCreate)
	grp.DELETE("/board/delete", board.ProcessBoardDelete)
	grp.PUT("/board/update", board.ProcessBoardUpdate)

	//random 명대사
	grp.GET("/lines/random", lines.ProcessRandomLine)

	// vote
	grp.PUT("/vote/update", vote.ProcessUpdateVote)
}
