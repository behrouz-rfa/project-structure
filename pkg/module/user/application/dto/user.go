package dto

type UserCreateInput struct {
	Email          string  `json:"email"`
	FirstName      string  `json:"firstName"`
	LastName       string  `json:"lastName"`
	MiddleName     *string `json:"middleName"`
	ProfilePicture *string `json:"profilePicture"`
	PhoneNumber    *string `json:"phoneNumber"`
	EmailVerified  *bool   `json:"emailVerified"`
	Active         *bool   `json:"active"`
	Role           *string `json:"role"`
}

type UserUpdateInput struct {
	Email          *string `json:"email"`
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName"`
	MiddleName     *string `json:"middleName"`
	ProfilePicture *string `json:"profilePicture"`
	PhoneNumber    *string `json:"phoneNumber"`
	EmailVerified  *bool   `json:"emailVerified"`
	Active         *bool   `json:"active"`
	Role           *string `json:"role"`
}
