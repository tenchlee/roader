package endpoint

import (
	"net"
	"sync"
	"errors"
	"fmt"
)

type ClientEndpoint struct {
	KcpEndpoint
	tcp_conn net.Conn
	buff_pool sync.Pool
}

func NewClientEndpoint() (client_ep *ClientEndpoint) {
	fmt.Println("new client endpoint")
	client_ep = new(ClientEndpoint)
	client_ep.buff_pool.New = func() interface{} {
		return make([]byte, 10240)
	}
	return
}

func (client_ep *ClientEndpoint) Init(tcp_conn net.Conn, originalIp string, originalPort uint16) {
	client_ep.tcp_conn = tcp_conn
	client_ep.status = EndpointInit

	active_ep := NewActiveEndpoint(client_ep)
	client_ep.peer = active_ep
	active_ep.Connect(originalIp, originalPort)
	go active_ep.Run()

}

func (client_ep *ClientEndpoint) Start() {
	var err error
	buff := make([]byte, 10240)
	var n int
	if client_ep.status != EndpointInit {
			err = errors.New("client endpoint status error")
			goto fini
	}
	client_ep.status = EndpointStart
	for {
		n, err = client_ep.tcp_conn.Read(buff)
		if err != nil {
			fmt.Println("client endpoint read fail:", err.Error())
			err = errors.New("client endpoint read fail: " + err.Error())
			break
		}
		fmt.Println("client endpoint recv len", n)
		err = client_ep.peer.Send(buff[:n])
		if err != nil {
			// todo error
			break
		}
	}
	fini:
	if err != nil {
		fmt.Println("err", err.Error())
	}
	client_ep.peer.Close()
	client_ep.Close()
	return

}

func (client_ep *ClientEndpoint) Send(data []byte) (err error) {
	_, err = client_ep.tcp_conn.Write(data)
	return
}

func (client_ep *ClientEndpoint) Close() {
	if client_ep.status == EndpointStop {
		return
	}
	fmt.Println("client ep close")
	client_ep.status = EndpointStop
	if client_ep.tcp_conn != nil {
		client_ep.tcp_conn.Close()
	}
	client_ep.peer = nil
}
