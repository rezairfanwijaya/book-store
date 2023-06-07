package admin

type responseAdminLogin struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

func FormatterResponseAdmin(admin Admin, token string) *responseAdminLogin {
	return &responseAdminLogin{
		ID:    admin.ID,
		Email: admin.Email,
		Token: token,
		Role:  admin.Role,
	}
}
