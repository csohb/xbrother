package xdb

import (
	"fmt"
	"testing"
	"xbrother/databases/mysql"
)

func TestGetLine(t *testing.T) {
	if err := mysql.ConnectMysql(); err != nil {
		panic(err)
	}

	db := mysql.Engine

	/*mem := entities.TBMember{
		Model: gorm.Model{},
		Name:  "철수",
		Score: 0,
	}

	db.Model(&entities.TBMember{}).Create(&mem)
	*/
	/*line := entities.TBLines{
		Model:      gorm.Model{},
		Line:       "형은 형이지. 형이 누구 꺼가 되고 말고가 어딨어요.",
		MemberCode: 3,
	}

	db.Model(&entities.TBLines{}).Create(&line)*/

	if lines, err := GetRandomLine(db); err != nil {
		panic(err)
	} else {
		fmt.Println(lines)
		fmt.Println(len(lines))
	}
}
