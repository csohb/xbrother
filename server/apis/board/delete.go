package board

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type BDeleteRequest struct {
	Seq uint   `json:"id"`
	Pwd string `json:"pwd"`
}

type ServiceBoardDelete struct {
	*context.Context
	req BDeleteRequest
}

func (app *ServiceBoardDelete) Service() *api_context.CommonResponse {
	logrus.Printf("path param id : %+v", app.req.Seq)
	logrus.Printf("json body pwd : %+v", app.req.Pwd)

	if pwd, err := xdb.GetMsgPwd(app.DB, app.req.Seq); err != nil {
		logrus.WithError(err).Error("select pwd failed.")
		return api_context.InternalServerErrJSON("db select err")
	} else {
		if pwd != app.req.Pwd {
			logrus.WithError(err).Error("not correct pwd.")
			return api_context.BadJSON("비밀번호가 일치하지 않습니다.")
		} else {
			if err = xdb.DeleteBoardMsg(app.DB, app.req.Seq); err != nil {
				logrus.WithError(err).Error("delete msg failed")
				return api_context.InternalServerErrJSON("db 삭제에 실패하였습니다.")
			}
		}
	}

	return api_context.EmptyJSON()
}

func (app *ServiceBoardDelete) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceBoardDelete) ApiName() interface{} {
	return "board-delete"
}

func ProcessBoardDelete(c echo.Context) error {
	return api_context.Process(&ServiceBoardDelete{Context: c.(*context.Context)})
}
