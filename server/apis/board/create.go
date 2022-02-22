package board

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type BCreateRequest struct {
	Nickname string `json:"nickname"`
	Pwd      string `json:"pwd"`
	Msg      string `json:"msg"`
}

type ServiceBoardCreate struct {
	*context.Context
	req BCreateRequest
}

func (app *ServiceBoardCreate) Service() *api_context.CommonResponse {

	if err := xdb.CreateBoardMsg(app.DB, app.req.Nickname, app.req.Pwd, app.req.Msg); err != nil {
		logrus.WithError(err).Error("insert db failed.")
		return api_context.InternalServerErrJSON("db insert 오류")
	}

	return api_context.EmptyJSON()
}

func (app *ServiceBoardCreate) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceBoardCreate) ApiName() interface{} {
	return "board-create"
}

func ProcessBoardCreate(c echo.Context) error {
	return api_context.Process(&ServiceBoardCreate{Context: c.(*context.Context)})
}
