package vote

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type List struct {
	ID        int    `json:"id"`
	Line      string `json:"line"`
	VoteCount int    `json:"vote_count"`
}

type LinesListDescResponse struct {
	list []List `json:"list"`
}

type ServiceLinesList struct {
	*context.Context
}

func (app *ServiceLinesList) Service() *api_context.CommonResponse {
	list, err := xdb.LinesListDesc(app.DB)
	if err != nil {
		logrus.Errorf("get vote list desc : %s", err)
		return api_context.InternalServerErrJSON("db err")
	}

	resp := LinesListDescResponse{}
	resp.list = make([]List, len(list))
	for i, v := range list {
		resp.list[i] = List{
			ID:        int(v.ID),
			Line:      v.Line,
			VoteCount: v.VoteCount,
		}
	}

	return api_context.SuccessJSON(&resp.list)
}

func (app *ServiceLinesList) GetRequestData() interface{} {
	return nil
}

func (app ServiceLinesList) ApiName() interface{} {
	return "vote-list"
}

func ProcessLinesListDesc(c echo.Context) error {
	return api_context.Process(&ServiceLinesList{
		Context: c.(*context.Context),
	})
}
