package main

import (
	tiktokapi "github.com/PICOF/simple-tiktok/kitex_gen/comment/commentservice"
	"log"
)

func main() {
	svr := tiktokapi.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
