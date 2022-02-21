package xdb

import (
	"gorm.io/gorm"
	"xbrother/databases/entities"
)

type BoardVo struct {
	gorm.Model
	Nickname string
	Pwd      string
	Msg      string
}

func (bl BoardVo) TableName() string {
	return entities.TBNameBoard
}

func GetBoardList(db *gorm.DB, page, count int) (ret []BoardVo, err error) {
	err = db.Model(&BoardVo{}).Offset(page * count).Limit(count).Find(&ret).Error
	if err != nil {
		return nil, err
	} else {
		return ret, nil
	}
}
