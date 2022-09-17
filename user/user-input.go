package user

type RegisterUserInput struct {
	Name       string `json:"name" form:"name" validation:"required"`
	Occupation string `json:"occupation" form:"occupation" validation:"required"`
	Email      string `json:"email" form:"email" validation:"required"`
	Password   string `json:"password" form:"password" validation:"required"`
}
