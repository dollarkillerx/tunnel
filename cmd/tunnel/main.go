package main

import (
	"github.com/dollarkillerx/tunnel/internal/conf"
	"github.com/dollarkillerx/tunnel/internal/server"

	"log"
)

func main() {
	conf.InitConf()

	ser := server.NewServer()
	log.Println("Init Server ...")
	if err := ser.Run(); err != nil {
		log.Fatalln(err)
	}
}
