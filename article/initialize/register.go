package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
)

// InitRegistry to init consul
func InitRegistry() (registry.Registry, *registry.Info) {
	r, err := consul.NewConsulRegister("127.0.0.1:8500",
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "7s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "15s",
		}))
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	// Using snowflake to generate service name.
	sf, err := snowflake.NewNode(3)
	if err != nil {
		klog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: "article_srv",
		Addr:        utils.NewNetAddr("tcp", "127.0.0.1:8882"),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
