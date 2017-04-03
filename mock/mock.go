// Testing for podcast-host

package mock

import (
	"codenex.us/ralph/podcast-host"
)

type UserService struct {
	UserFn      func(id int64) (*podhost.User, error)
	UserInvoked bool

	CreateUserFn      func(*podhost.User) error
	CreateuserInvoked bool

	DeleteUserFn      func(id int64) error
	DeleteUserInvoked bool
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int64) (*podhost.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}

// CreateUser invokes the mock implementation and marks the function as invoked.
func (s *UserService) CreateUser(u *podhost.User) error {
	s.CreateuserInvoked = true
	return s.CreateUserFn(u)
}

// DeleteUser invokes the mock implementation and marks the function as invoked.
func (s *UserService) DeleteUser(id int64) error {
	s.DeleteUserInvoked = true
	return s.DeleteUserFn(id)
}
