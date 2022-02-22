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

func CreateBoardMsg(db *gorm.DB, nickname, pwd, msg string) error {
	if err := db.Model(&BoardVo{}).Create(&BoardVo{
		Nickname: nickname,
		Pwd:      pwd,
		Msg:      msg,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GetMsgPwd(db *gorm.DB, seq uint) (pwd string, err error) {
	if err = db.Model(&BoardVo{}).Select("pwd").Where("id=?", seq).Take(&pwd).Error; err != nil {
		return "", err
	} else {
		return pwd, nil
	}
}

func DeleteBoardMsg(db *gorm.DB, seq uint) error {
	if err := db.Delete(&BoardVo{}, seq).Error; err != nil {
		return err
	}
	return nil
}
