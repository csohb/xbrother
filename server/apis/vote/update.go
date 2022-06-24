package vote

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type UpdateVoteRequest struct {
	ID uint `query:"id"`
}

type ServiceUpdateVote struct {
	*context.Context
	req UpdateVoteRequest
}

func (app *ServiceUpdateVote) Service() *api_context.CommonResponse {
	if err := xdb.UpdateVote(app.DB, app.req.ID); err != nil {
		logrus.WithError(err).Error("update vote count failed.")
		return api_context.InternalServerErrJSON("투표에 실패하였습니다.")
	}
	return api_context.EmptyJSON()
}

func (app *ServiceUpdateVote) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateVote) ApiName() interface{} {
	return "vote-update"
}

func ProcessUpdateVote(c echo.Context) error {
	return api_context.Process(&ServiceUpdateVote{Context: c.(*context.Context)})
}
