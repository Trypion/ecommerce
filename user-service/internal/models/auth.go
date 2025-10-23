package models

type AuthUser struct {
	ID    string
	Email string
	Name  string
	Role  string
}

type AuthLogin struct {
	User  AuthUser
	Token string
}
