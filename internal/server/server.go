package server

import (
	"github.com/dollarkillerx/tunnel/internal/conf"

	"log"
	"net"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	for _, v := range conf.Conf.ProxyItem {
		dial, err := net.Listen("tcp", v.ExternalAddress)
		if err != nil {
			return err
		}

		proxy := NewProxy(dial, v.ProxyAddress)
		go func() {
			err := proxy.Run()
			if err != nil {
				log.Println(err)
			}
		}()
	}

	return nil
}
