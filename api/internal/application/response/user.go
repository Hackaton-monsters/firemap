package response

type LoginUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type RegisteredUser struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
