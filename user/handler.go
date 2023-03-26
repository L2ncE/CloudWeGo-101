package main

import (
	"context"
	user "github.com/L2ncE/CloudWeGo-101/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.ResgisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetArticleNum implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetArticleNum(ctx context.Context, req *user.GetArticleNumRequest) (resp *user.GetArticleNumResponse, err error) {
	// TODO: Your code here...
	return
}

// AddArticleNum implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddArticleNum(ctx context.Context, req *user.AddArticleNumRequest) (resp *user.AddArticleNumResponse, err error) {
	// TODO: Your code here...
	return
}
