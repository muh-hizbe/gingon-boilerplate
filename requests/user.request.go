package requests

type UserRequest struct {
	Name  string `json:"name" form:"name" gorm:"required"`
	Email string `json:"email" form:"name" gorm:"required"`
}
