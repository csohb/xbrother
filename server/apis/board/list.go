package board

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type BListRequest struct {
	Page  int `query:"page"`
	Count int `query:"count"`
}

type BData struct {
	Seq         uint   `json:"seq"`
	Text        string `json:"text"`
	Nickname    string `json:"nickname"`
	RegDatetime string `json:"reg_datetime"`
}

type BListResponse struct {
	List  []BData `json:"list"`
	Total int     `json:"total"`
}

type ServiceBoardList struct {
	*context.Context
	req BListRequest
}

func (app *ServiceBoardList) Service() *api_context.CommonResponse {
	resp := BListResponse{}

	if list, err := xdb.GetBoardList(app.DB, app.req.Page-1, app.req.Count); err != nil {
		logrus.WithError(err).Error("select list failed : ", err)
		return api_context.InternalServerErrJSON("db 조회 실패")
	} else {
		resp.List = make([]BData, len(list))
		for i, v := range list {
			resp.List[i] = BData{
				Seq:         v.ID,
				Text:        v.Msg,
				Nickname:    v.Nickname,
				RegDatetime: "2005-" + v.CreatedAt.Format("01-02 15:04:05"),
			}
		}
		resp.Total = len(list)
	}

	return api_context.SuccessJSON(&resp)
}

func (app *ServiceBoardList) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceBoardList) ApiName() interface{} {
	return "board-list"
}

func ProcessBoardList(c echo.Context) error {
	return api_context.Process(&ServiceBoardList{Context: c.(*context.Context)})
}
