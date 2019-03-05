package gate_listener

import (
	"net"
	"github.com/tenchlee/roader/endpoint"
	"syscall"
	"fmt"
)

const SO_ORIGINAL_DST = 80

type TcpGateListener struct {
	GateListener
}

func NewTcpGateListener() (tcp_gl *TcpGateListener) {
	tcp_gl = new (TcpGateListener)
	return
}

func (tcp_gl *TcpGateListener) Listen(port int) (err error) {
	tcp_gl.port = port
	//net_cfg := net.ListenConfig{
	//	Control:   nil,
	//}
	addr, err := net.ResolveTCPAddr("tcp", tcp_gl.getAddrString())
	if err != nil {
		return
	}
	conn, err := net.ListenTCP("tcp", addr)
	//conn, err := net_cfg.Listen(context.Background(), "tcp", tcp_gl.getAddrString())
	if err != nil {
		return err
	}
	for {
		new_conn, err := conn.AcceptTCP()
		if err != nil {
			return err
		}
		raw_conn, err := new_conn.SyscallConn()
		if err != nil {
			return err
		}
		raw_conn.Control(func(fd uintptr) {
			// setsockopt
			addr, err := syscall.GetsockoptIPv6Mreq(int(fd), syscall.IPPROTO_IP, SO_ORIGINAL_DST)
			if err != nil {
				fmt.Println("err get sockopt")
				return
			}
			ip_arr := addr.Multiaddr[4:8]
			original_ip :=  fmt.Sprintf("%d.%d.%d.%d", ip_arr[0], ip_arr[1], ip_arr[2], ip_arr[3])
			original_port := uint16(addr.Multiaddr[2])<<8 + uint16(addr.Multiaddr[3])
			fmt.Println(original_ip, original_port)
			go tcp_gl.NewConnection(new_conn, original_ip, original_port)
		})
	}
	return
}

func (tcp_gl *TcpGateListener) NewConnection(conn net.Conn, originalIp string, originalPort uint16) {
	client_ep := endpoint.NewClientEndpoint()
	client_ep.Init(conn, originalIp, originalPort)
	go client_ep.Start()
}