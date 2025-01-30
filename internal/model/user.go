package model

type UserInfo struct {
	Password    string `db:"password"`
	Email       string `db:"email"`
	Role        string `db:"role"`
}
