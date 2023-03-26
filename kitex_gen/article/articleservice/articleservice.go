// Code generated by Kitex v0.4.4. DO NOT EDIT.

package articleservice

import (
	"context"
	article "github.com/L2ncE/CloudWeGo-101/kitex_gen/article"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return articleServiceServiceInfo
}

var articleServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ArticleService"
	handlerType := (*article.ArticleService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PostArticle": kitex.NewMethodInfo(postArticleHandler, newArticleServicePostArticleArgs, newArticleServicePostArticleResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "article",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func postArticleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*article.ArticleServicePostArticleArgs)
	realResult := result.(*article.ArticleServicePostArticleResult)
	success, err := handler.(article.ArticleService).PostArticle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newArticleServicePostArticleArgs() interface{} {
	return article.NewArticleServicePostArticleArgs()
}

func newArticleServicePostArticleResult() interface{} {
	return article.NewArticleServicePostArticleResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PostArticle(ctx context.Context, req *article.PostArticleResponse) (r *article.PostArticleResponse, err error) {
	var _args article.ArticleServicePostArticleArgs
	_args.Req = req
	var _result article.ArticleServicePostArticleResult
	if err = p.c.Call(ctx, "PostArticle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
