package xdb

import "xbrother/databases/entities"

type MemberVo struct {
	ID   uint
	Name string
}

func (mv MemberVo) TableName() string {
	return entities.TBNameMember
}
