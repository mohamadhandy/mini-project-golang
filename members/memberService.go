package members

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type MemberService interface {
	RegisterMember(MemberRegisterInput) (Member, error)
	Login(LoginInput) (Member, error)
}

type serviceMember struct {
	memberRepositoryDB MemberRepositoryDB
}

func NewServiceMember(memberRepo MemberRepositoryDB) *serviceMember {
	return &serviceMember{memberRepositoryDB: memberRepo}
}

func (s *serviceMember) RegisterMember(input MemberRegisterInput) (Member, error) {
	member := Member{}
	member.Email = input.Email
	member.Role = input.Role
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return member, err
	}
	member.Password = string(passwordHash)
	newMember, err := s.memberRepositoryDB.RegisterMember(member)
	if err != nil {
		return newMember, err
	}
	return newMember, nil
}

func (s *serviceMember) Login(input LoginInput) (Member, error) {
	email := input.Email
	password := input.Password

	member, err := s.memberRepositoryDB.FindByEmail(email)
	if err != nil {
		return member, err
	}
	if member.Email == "" {
		return member, errors.New("no member on that email found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password))
	if err != nil {
		return member, err
	}
	return member, nil
}
