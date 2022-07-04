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
	if err := db.Model(&VoteVo{}).Where("id = ?", lineId).Updates(map[string]interface{}{
		"vote_count": gorm.Expr("vote_count + ?", 1),
	}).Error; err != nil {
		return err
	}
	return nil
}

func LinesListDesc(db *gorm.DB) (ret []VoteVo, err error) {
	if err = db.Model(&VoteVo{}).Order("vote_count desc").Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
