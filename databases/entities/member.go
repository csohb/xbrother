package entities

import "gorm.io/gorm"

const TBNameMember = "member"

type TBMember struct {
	gorm.Model
	Name  string `gorm:"varchar(16)"`
	Score int
	Lines []TBLines `gorm:"foreignKey:member_code;references:id"`
}

func (t TBMember) TableName() string {
	return TBNameMember
}
