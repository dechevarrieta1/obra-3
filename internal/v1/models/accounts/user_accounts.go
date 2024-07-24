package accountsmodelsv1

type AccountUserRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AccountUserResponse struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	JWT      string `json:"jwt"`
}

type AccountUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
