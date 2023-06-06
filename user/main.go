package main

import (
	user "github.com/L2ncE/CloudWeGo-101/kitex_gen/user/userservice"
	"github.com/L2ncE/CloudWeGo-101/user/initialize"
	mysql "github.com/L2ncE/CloudWeGo-101/user/pkg"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry()
	db := initialize.InitDB()

	// Create new server.
	srv := user.NewServer(&UserServiceImpl{
		MySqlManager: mysql.NewUserManager(db),
	},
		server.WithServiceAddr(utils.NewNetAddr("tcp", "10.20.72.188:8881")),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user_srv"}),
	)
	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
