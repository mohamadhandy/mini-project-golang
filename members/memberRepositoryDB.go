package members

import (
	"miniprojectgo/logger"

	"gorm.io/gorm"
)

type MemberRepositoryDB interface {
	RegisterMember(member Member) (Member, error)
}

type memberRepositoryDB struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepositoryDB {
	return &memberRepositoryDB{db}
}

func (s *memberRepositoryDB) RegisterMember(member Member) (Member, error) {
	var err error
	if err = s.db.Create(&member).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return member, nil
	}
	return member, nil
}
