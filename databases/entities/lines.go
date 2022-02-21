package entities

import "gorm.io/gorm"

const TBNameLines = "lines"

type TBLines struct {
	gorm.Model
	Line       string `gorm:"type:varchar(4000)"`
	MemberCode uint
	Member     TBMember `gorm:"foreignKey:member_code"`
}

func (t TBLines) TableName() string {
	return TBNameLines
}
