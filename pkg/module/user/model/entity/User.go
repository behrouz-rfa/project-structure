package entity

import (
	"time"
)

type User struct {
	ID             string
	Email          string
	FirstName      string
	LastName       string
	MiddleName     *string
	ProfilePicture *string
	PhoneNumber    string
	EmailVerified  bool
	Active         bool
	Role           Role
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Role string

const (
	RoleAdmin          Role = "admin"
	RoleContentManager Role = "content_manager"
	RoleStudent        Role = "student"
	RoleTeacher        Role = "teacher"
	RoleSchoolManager  Role = "school_manager"
	RoleParent         Role = "parent"
	RoleSuperAdmin     Role = "super_admin"
	RoleGuest          Role = "guest"
)

func (u User) GetID() string {
	return u.ID
}

func (u User) SetID(id string) {

}

func (u User) GenerateID() {

}

func (u User) SetCreatedAt() {

}

func (u User) SetUpdatedAt() {

}

func (u User) CollectionName() string {
	return "users"
}
