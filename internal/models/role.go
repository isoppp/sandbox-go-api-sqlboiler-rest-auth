package models

type UserRole string

const (
	RoleAdmin = UserRole("Admin")
	RoleUser  = UserRole("User")
)
