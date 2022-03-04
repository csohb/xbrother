package xdb

import (
	"gorm.io/gorm"
	"xbrother/databases/entities"
)

type RandomLinesVo struct {
	ID         uint
	Line       string
	MemberCode uint
	Member     MemberVo `gorm:"foreignKey:member_code"`
}

func (rv RandomLinesVo) TableName() string {
	return entities.TBNameLines
}

func GetRandomLine(db *gorm.DB) (ret []RandomLinesVo, err error) {
	if err = db.Model(&RandomLinesVo{}).Joins("Member").Find(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}
