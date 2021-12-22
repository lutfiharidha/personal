package dtos

type AuthCreateDTO struct {
	ID       string `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name" binding:"required"`
	Image    string `json:"image" form:"image"`
}
type AuthUpdateDTO struct {
	ID       string `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password"  binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Image    string `json:"image" form:"image"`
}

type AuthLoginDTO struct {
	Email string `json:"email" form:"email" binding:"required"`
	// Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password"`
}

type AuthRestoreDTO struct {
	ID string `json:"id" form:"id"`
}
