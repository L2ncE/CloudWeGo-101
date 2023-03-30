package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"

	"github.com/L2ncE/CloudWeGo-101/kitex_gen/article"
)

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct {
	MySqlManager
	UserManager
}

// UserManager defines the ACL(Anti Corruption Layer)
type UserManager interface {
	ArticleNumPlusOne(uid int64) error
}

// MySqlManager implements database operations.
type MySqlManager interface {
	Post(title, content string, uid int64) error
}

// PostArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) PostArticle(_ context.Context, req *article.PostArticleRequest) (resp *article.PostArticleResponse, err error) {
	err = s.MySqlManager.Post(req.Title, req.Content, req.Uid)
	if err != nil {
		klog.Error("post article error, err:", err)
		return nil, status.Err(codes.Internal, "post error")
	}
	err = s.UserManager.ArticleNumPlusOne(req.Uid)
	if err != nil {
		klog.Error("article num plus one error, err:", err)
		return nil, status.Err(codes.Internal, "post error")
	}
	return nil, nil
}
