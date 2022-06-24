package entities

import "gorm.io/gorm"

const TBNameLines = "lines"

type TBLines struct {
	gorm.Model
	Line       string `gorm:"type:varchar(4000)"`
	MemberCode uint
	VoteCount  int
}

func (t TBLines) TableName() string {
	return TBNameLines
}
