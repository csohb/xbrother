package entities

import "gorm.io/gorm"

const TBNameMember = "member"

type TBMember struct {
	gorm.Model
	Name  string `gorm:"varchar(16)"`
	Score int
}

func (t TBMember) TableName() string {
	return TBNameMember
}
