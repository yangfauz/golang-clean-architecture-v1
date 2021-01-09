package user

type GetUserDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateUserInput struct {
	Name string `form:"name" binding:"required"`
	Role string `form:"role" binding:"required"`
}

type UpdateUserInput struct {
	Name string `json:"name" binding:"required"`
	Role string `json:"role" binding:"required"`
}
