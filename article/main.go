package main

import (
	article "github.com/L2ncE/CloudWeGo-101/kitex_gen/article/articleservice"
	"log"
)

func main() {
	svr := article.NewServer(new(ArticleServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
