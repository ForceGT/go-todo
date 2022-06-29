package model

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	RoleID    int    `json:"roleId"`
}
