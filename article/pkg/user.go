package pkg

import (
	"context"

	"github.com/L2ncE/CloudWeGo-101/kitex_gen/user"
	"github.com/L2ncE/CloudWeGo-101/kitex_gen/user/userservice"
)

type UserManager struct {
	UserClient userservice.Client
}

func (u *UserManager) ArticleNumPlusOne(uid int64) error {
	_, err := u.UserClient.AddArticleNum(context.Background(), &user.AddArticleNumRequest{UserId: uid})
	return err
}
