package main

import (
	"context"
	article "github.com/L2ncE/CloudWeGo-101/kitex_gen/article"
)

// ArticleServiceImpl implements the last service interface defined in the IDL.
type ArticleServiceImpl struct{}

// PostArticle implements the ArticleServiceImpl interface.
func (s *ArticleServiceImpl) PostArticle(ctx context.Context, req *article.PostArticleResponse) (resp *article.PostArticleResponse, err error) {
	// TODO: Your code here...
	return
}
