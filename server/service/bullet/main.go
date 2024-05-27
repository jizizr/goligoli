package main

import (
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"log"
)

func main() {
	svr := bullet.NewServer(new(BulletServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
