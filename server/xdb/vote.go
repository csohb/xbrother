package xdb

import (
	"gorm.io/gorm"
	"xbrother/databases/entities"
)

type VoteVo struct {
	ID        uint
	Line      string
	VoteCount int
}

func (vv VoteVo) TableName() string {
	return entities.TBNameLines
}

func UpdateVote(db *gorm.DB, lineId uint) error {
	if err := db.Model(&VoteVo{}).Where("id = ? ", lineId).
		UpdateColumn("vote_count", gorm.Expr("column + ?", 1)).Error;
		err != nil {
		return err
	}
	return nil
}
