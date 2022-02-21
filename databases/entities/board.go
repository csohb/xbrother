package entities

import "gorm.io/gorm"

const TBNameBoard = "board"

type TBBoard struct {
	gorm.Model
	Nickname string `gorm:"index;not null"`
	Pwd      string `gorm:"type:varchar(100)"`
	Msg      string `gorm:"type:varchar(4000)"`
}

func (t TBBoard) TableName() string {
	return TBNameBoard
}
