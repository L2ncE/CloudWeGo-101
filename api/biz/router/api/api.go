// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	api "github.com/L2ncE/CloudWeGo-101/api/biz/handler/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/article", append(_post_rticleMw(), api.PostArticle)...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/article", append(_get_rticlenumMw(), api.GetArticleNum)...)
		_user.POST("/article", append(__dd_rticlenumMw(), api.AddArticleNum)...)
		{
			_login := _user.Group("/login", _loginMw()...)
			_login.POST("/", append(_login0Mw(), api.Login)...)
		}
		{
			_register := _user.Group("/register", _registerMw()...)
			_register.POST("/", append(_register0Mw(), api.Register)...)
		}
	}
}
