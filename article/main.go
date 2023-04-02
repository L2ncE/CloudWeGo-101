package main

import (
	"github.com/L2ncE/CloudWeGo-101/article/initialize"
	"github.com/L2ncE/CloudWeGo-101/article/pkg"
	article "github.com/L2ncE/CloudWeGo-101/kitex_gen/article/articleservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry()
	db := initialize.InitDB()

	// Create new server.
	srv := article.NewServer(&ArticleServiceImpl{
		MySqlManager: pkg.NewArticleManager(db),
	},
		server.WithServiceAddr(utils.NewNetAddr("tcp", "127.0.0.1:8882")),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "article_srv"}),
	)
	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
