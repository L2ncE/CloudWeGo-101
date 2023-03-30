package global

import (
	"github.com/L2ncE/CloudWeGo-101/kitex_gen/article/articleservice"
	"github.com/L2ncE/CloudWeGo-101/kitex_gen/user/userservice"
)

var (
	GlobalUserClient    userservice.Client
	GlobalArticleClient articleservice.Client
)
