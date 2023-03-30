package init

import (
	user "github.com/L2ncE/CloudWeGo-101/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
)

// InitUser to init user service
func InitUser() user.Client {
	// init resolver
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatalf("new nacos client failed: %s", err.Error())
	}

	// create a new client
	c, err := user.NewClient(
		"user_srv",
		client.WithResolver(r), // service discovery
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user_srv"}),
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	return c
}
