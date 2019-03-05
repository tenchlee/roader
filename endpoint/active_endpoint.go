package endpoint

import (
	"fmt"
	"net"
	"github.com/tenchlee/roader/msg"
	"github.com/tenchlee/roader/binary"
	"github.com/golang/protobuf/proto"
)

type ActiveEndpoint struct {
	KcpEndpoint

	addr *net.UDPAddr
}

func NewActiveEndpoint(peer IEndpoint) (active_ep *ActiveEndpoint) {
	active_ep = new (ActiveEndpoint)
	active_ep.peer = peer
	return
}

func (active_ep *ActiveEndpoint) Connect(originalIp string, originalPort uint16) (err error) {
	fmt.Println("active endpoint connect")
	active_ep.addr, err = net.ResolveUDPAddr("udp", "192.168.195.133:8000")
	if err != nil {
		fmt.Println("active endpoint connect error", err.Error())
		return err
	}
	active_ep.init(func(buf []byte, size int) {
		fmt.Println("active ep send", size)
		active_ep.conn.Conn.WriteTo(buf[:size], active_ep.addr)
	})
	msg := &rdp_msg.Rdp_Msg{
		Protocol: rdp_msg.Protocol_PROTOCOL_TCP.Enum(),
		Business: proto.String("x8ey1"),
		DstIp: proto.Uint32(binary.IPv4IPtoUint(originalIp)),
		DstPort: proto.Uint32(uint32(binary.Htons(originalPort))),
		UserId: proto.String("tench"),
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("proto marshal fail", err)
		return
	}
	active_ep.SendSYN(data)
	return nil
}

func (active_ep *ActiveEndpoint) Run() error{
	fmt.Println("active endpoint run")
	go active_ep.process()
	// send back
	for {
		select {
			case data := <- active_ep.user_recv_queue:
				active_ep.peer.Send(data.Bytes())
		}
	}
	return nil
}
