package rpc

import (
	article "github.com/L2ncE/CloudWeGo-101/kitex_gen/article/articleservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
)

// initArticle to init user service
func initArticle() article.Client {
	// init resolver
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}

	// create a new client
	c, err := article.NewClient(
		"article_srv",
		client.WithResolver(r), // service discovery
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "article_srvå"}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	return c
}
