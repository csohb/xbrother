package board

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type BUpdateRequest struct {
	Seq uint   `json:"id"`
	Pwd string `json:"pwd"`
	Msg string `json:"msg"`
}

type ServiceBoardUpdate struct {
	*context.Context
	req BUpdateRequest
}

func (app *ServiceBoardUpdate) Service() *api_context.CommonResponse {
	pwd, err := xdb.GetMsgPwd(app.DB, app.req.Seq)
	if err != nil {
		logrus.WithError(err).Error("get pwd err")
		return api_context.InternalServerErrJSON("비밀번호를 불러오는데 실패하였습니다.")
	}
	if pwd != app.req.Pwd {
		logrus.WithError(err).Error("wrong pwd")
		return api_context.BadJSON("비민번호가 일치하지 않습니다.")
	}

	if err = xdb.UpdateBoardMsg(app.DB, app.req.Seq, app.req.Msg); err != nil {
		logrus.WithError(err).Error("update board failed.")
		return api_context.InternalServerErrJSON("게시글 수정에 실패하였습니다.")
	}

	return api_context.EmptyJSON()
}

func (app *ServiceBoardUpdate) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceBoardUpdate) ApiName() interface{} {
	return "board-update"
}

func ProcessBoardUpdate(c echo.Context) error {
	return api_context.Process(&ServiceBoardUpdate{Context: c.(*context.Context)})
}
