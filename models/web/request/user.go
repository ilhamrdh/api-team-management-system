package request

type UserRequest struct {
	FirstName string `validate:"required,min=1,max=100" json:"first_name"`
	LastName  string `validate:"required,min=1,max=100" json:"last_name"`
	Email     string `validate:"required,min=1,max=100" json:"email"`
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Password  string `validate:"required,min=1,max=100" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}

type UserDetail struct {
	Username  string `validate:"required,min=1,max=100" json:"username"`
	Authorize string `json:"authorized"`
}
