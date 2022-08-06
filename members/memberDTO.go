package members

type MemberDTO struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FormatMemberDTO(member Member, token string) MemberDTO {
	memberDTO := MemberDTO{
		ID:    member.ID,
		Email: member.Email,
		Role:  member.Role,
		Token: token,
	}
	return memberDTO
}
