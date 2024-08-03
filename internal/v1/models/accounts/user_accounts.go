package accountsmodelsv1

type AccountUserRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	AccountID string `json:"account_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
type AccountUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
