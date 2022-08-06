package members

import (
	"fmt"
	"miniprojectgo/logger"

	"gorm.io/gorm"
)

type MemberRepositoryDB interface {
	RegisterMember(Member) (Member, error)
	FindByEmail(string) (Member, error)
	FindById(id int) (Member, error)
}

type memberRepositoryDB struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepositoryDB {
	return &memberRepositoryDB{db}
}

func (s *memberRepositoryDB) RegisterMember(member Member) (Member, error) {
	fmt.Println("member", member)
	var err error
	if err = s.db.Create(&member).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return member, nil
	}
	return member, nil
}

func (s *memberRepositoryDB) FindByEmail(email string) (Member, error) {
	var member Member
	var err error
	if err = s.db.Where("email = ?", email).Find(&member).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return member, err
	}
	return member, nil
}

func (s *memberRepositoryDB) FindById(id int) (Member, error) {
	var member Member
	var err error
	if err = s.db.Where("member_id = ?", id).Find(&member).Error; err != nil {
		logger.Error("Unexpected Error: " + err.Error())
		return member, err
	}
	return member, nil
}
