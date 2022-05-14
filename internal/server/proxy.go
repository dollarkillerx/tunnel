package server

import (
	"github.com/dollarkillerx/tunnel/internal/utils"

	"log"
	"net"
)

type Proxy struct {
	local        net.Listener
	proxyAddress string
}

func NewProxy(local net.Listener, proxyAddress string) *Proxy {
	return &Proxy{
		local:        local,
		proxyAddress: proxyAddress,
	}
}

func (p *Proxy) Run() error {
	for {
		accept, err := p.local.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go p.proxy(accept)
	}
}

func (p *Proxy) proxy(conn net.Conn) {
	proxyConn, err := net.Dial("tcp", p.proxyAddress)
	if err != nil {
		log.Println(err)
		return
	}

	err = utils.Transport(proxyConn, conn)
	if err != nil {
		log.Println(err)
		return
	}
}
