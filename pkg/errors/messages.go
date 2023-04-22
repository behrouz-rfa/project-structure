package errors

const (
	ErrCanNotCreateUser              = "could not create user"
	ErrCanNotGetUser                 = "could not get user"
	ErrCanNotGetUserAfterCreate      = "could not get user after create"
	ErrCanNotGetUserAfterUpdate      = "could not get user after update"
	ErrCanNotUpdateUser              = "could not  update user"
	ErrRoleIsNotValid                = "role is not valid"
	ErrUserNotFoundForAddClassMember = "user not found for add classroom member"
	ErrClassRoomNotFound             = "classroom not found"
	ErrCCouldNotRemoveClassroom      = "could not remove classroom"
	ErrCouldNotAddMember             = "could not add member"
	ErrUserAlreadyExistInClassroom   = "user already exists in classroom"
	ErrSchoolNotFound                = "school not found"
	ErrUpdateSchool                  = "failed to update school"
	ErrCouldNotCreateClassroom       = "could not create classroom"
	ErrCouldNotUpdateClassroom       = "could not update classroom"
)
