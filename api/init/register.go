package init

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

// InitRegistry to init consul
func InitRegistry() (registry.Registry, *registry.Info) {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}

	r := consul.NewConsulRegister(consulClient,
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "7s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "15s",
		}))

	// Using snowflake to generate service name.
	sf, err := snowflake.NewNode(3)
	if err != nil {
		hlog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: "api",
		Addr:        utils.NewNetAddr("tcp", "127.0.0.1:8882"),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
