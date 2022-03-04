package lines

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"xbrother/common/api_context"
	"xbrother/server/apis/context"
	"xbrother/server/xdb"
)

type Line struct {
	ID         uint     `json:"id"`
	Line       string   `json:"line"`
	MemberCode uint     `json:"member_code"`
	Member     MemberVo `json:"member"`
}

type MemberVo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type RandomLineResponse struct {
	Line []Line `json:"line"`
}

type ServiceRandomLine struct {
	*context.Context
}

func (app *ServiceRandomLine) Service() *api_context.CommonResponse {
	resp := RandomLineResponse{}

	if ret, err := xdb.GetRandomLine(app.DB); err != nil {
		logrus.WithError(err).Error("get random line failed.")
		return api_context.InternalServerErrJSON("랜덤 명대사 불러오기에 실패하였습니다.")
	} else {
		logrus.Printf("len(ret) : %+v", len(ret))
		logrus.Printf("random int : %+v", math.Round(rand.Float64()*float64(len(ret))))
		//randomInt := math.Round(rand.Float64() * float64(len(ret)))

		resp.Line = make([]Line, len(ret))
		for i, v := range ret {
			logrus.Printf(" i : %+v", i)
			resp.Line[i] = Line{
				ID:         v.ID,
				Line:       v.Line,
				MemberCode: v.MemberCode,
				Member: MemberVo{
					ID:   v.Member.ID,
					Name: v.Member.Name,
				},
			}
		}
		//logrus.Printf("ret : %+v\n", resp.Line[int(randomInt)])
	}
	return api_context.SuccessJSON(resp)
}

func (app *ServiceRandomLine) GetRequestData() interface{} {
	return nil
}

func (app ServiceRandomLine) ApiName() interface{} {
	return "lines-random"
}

func ProcessRandomLine(c echo.Context) error {
	return api_context.Process(&ServiceRandomLine{Context: c.(*context.Context)})
}
