package request

type CreateAdminSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateSignUpMentorRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Avatar   string `json:"avatar"`
}

type CreateSignInMentorRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
