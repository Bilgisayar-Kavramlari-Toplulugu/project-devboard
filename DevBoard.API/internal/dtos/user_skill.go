package dtos

type UserSkillDTO struct {
	Id        int64  `json:"id"`
	SkillId   int    `json:"skillId"`
	SkillName string `json:"skillName"`
}
