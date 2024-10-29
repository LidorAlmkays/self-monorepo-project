package userResponseManager

type UserResponseManager interface {
	UserAdded(userId int) error
}
