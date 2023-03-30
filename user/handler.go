package main

import (
	"context"

	"github.com/L2ncE/CloudWeGo-101/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MySqlManager
}

// MySqlManager implements database operations.
type MySqlManager interface {
	LoginCheck(username, password string) (bool, int64, error)
	CreateUser(username, password string) error
	GetArticleNum(uid int64) (int64, error)
	ArticleNumPlusOne(uid int64) error
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(_ context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	flag, uid, err := s.MySqlManager.LoginCheck(req.Username, req.Password)
	if err != nil {
		klog.Error("login check error, err :", err)
		return nil, status.Err(codes.Internal, "login error")
	}
	if !flag {
		klog.Info("wrong password")
		return nil, status.Err(codes.Internal, "login error")
	}
	return &user.LoginResponse{Uid: uid}, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(_ context.Context, req *user.RegisterRequest) (resp *user.ResgisterResponse, err error) {
	err = s.MySqlManager.CreateUser(req.Username, req.Password)
	if err != nil {
		klog.Error("register error, err :", err)
		return nil, status.Err(codes.Internal, "register error")
	}
	return nil, nil
}

// GetArticleNum implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetArticleNum(_ context.Context, req *user.GetArticleNumRequest) (resp *user.GetArticleNumResponse, err error) {
	count, err := s.MySqlManager.GetArticleNum(req.UserId)
	if err != nil {
		klog.Error("get article num error, err :", err)
		return nil, status.Err(codes.Internal, "get article num error")
	}
	return &user.GetArticleNumResponse{Num: count}, nil
}

// AddArticleNum implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddArticleNum(_ context.Context, req *user.AddArticleNumRequest) (resp *user.AddArticleNumResponse, err error) {
	err = s.MySqlManager.ArticleNumPlusOne(req.UserId)
	if err != nil {
		klog.Error("article num plus one error, err :", err)
		return nil, status.Err(codes.Internal, "add article num error")
	}
	return nil, nil
}
