package members

import "golang.org/x/crypto/bcrypt"

type MemberService interface {
	RegisterMember(MemberRegisterInput) (Member, error)
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
