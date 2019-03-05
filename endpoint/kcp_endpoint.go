package endpoint

import (
	"time"
	"fmt"
	"github.com/tenchlee/gcp"
	"github.com/tenchlee/roader/conn_pool"
)

const (
	EndpointInit	 = 1
	EndpointStart	 = 2
	EndpointStop	 = 3
)

type IEndpoint interface {
	Send(data []byte) (err error)
	Close()
}

type KcpEndpoint struct {
	IEndpoint
	conv uint32
	peer IEndpoint
	status int
	kcp *gcp.KCP
	conn *conn_pool.GcpConn

	input_queue			chan []byte
	output_queue  		chan []byte
	user_recv_queue 	chan *gcp.ByteBuffer
}

func (ep *KcpEndpoint) init(output_cb func(buf []byte, size int)) {
	ep.conn, ep.conv = conn_pool.G_UdpConnPool.Get(ep)
	ep.input_queue = make (chan []byte, 10240)
	ep.output_queue = make (chan []byte, 10240)
	ep.user_recv_queue = make (chan *gcp.ByteBuffer, 10240)

	ep.kcp = gcp.NewKCP(ep.conv, output_cb)
}

func (ep *KcpEndpoint) process() {
	tick := time.Tick(10 * time.Millisecond)
	var data []byte
	for {
		select {
		case data = <- ep.input_queue:
			//fmt.Println("processUserData", len(data))
			start := uint32(time.Now().UnixNano()/1e6)
			i := 0
			for {
				ret := ep.kcp.Input(data, true, true)
				ep.conn.PutBuff(data)
				if ret != 0 {
					fmt.Println("input fail ret", ret)
				}
				i++
				if len (ep.input_queue) > 0 && i < 10 {
					data = <- ep.input_queue
				} else {
					break
				}
			}
			ep.kcp.ForceFlush()
			closed, _ := ep.kcp.Status()
			if closed {
				fmt.Println("closed")
				ep.conn.PutConn(ep.conv)
				return
			}
			end1 := uint32(time.Now().UnixNano()/1e6)
			ep.processUserData()
			end2 := uint32(time.Now().UnixNano()/1e6)
			if end1 - start > 10 || end2 - end1 > 10 {
				fmt.Println("input end1", end1 - start, "end2", end2 - end1)
			}
		case data = <- ep.output_queue:
			//fmt.Println("send", len(data))
			start := uint32(time.Now().UnixNano()/1e6)
			i := 0
			for {
				ep.kcp.Send(data)
				i++
				if len (ep.output_queue) > 0 && i < 10 {
					data = <- ep.output_queue
				} else {
					break
				}
			}
			ep.kcp.ForceFlush()
			closed, _ := ep.kcp.Status()
			if closed {
				break
			}
			end := uint32(time.Now().UnixNano()/1e6)
			if end - start > 10 {
				fmt.Println("kcp send", end-start)
			}
		case <- tick:
			start := uint32(time.Now().UnixNano()/1e6)
			ep.kcp.Update()
			closed, _ := ep.kcp.Status()
			if closed {
				fmt.Println("closed")
				ep.conn.PutConn(ep.conv)
				return
			}
			end1 := uint32(time.Now().UnixNano()/1e6)
			ep.processUserData()
			end2 := uint32(time.Now().UnixNano()/1e6)
			if end1 - start > 10 || end2 - end1 > 10 {
				fmt.Println("tock", end1 - start, "end2", end2 - end1)
			}
		}
	}
}

func (ep *KcpEndpoint) processUserData() {
	if  len(ep.user_recv_queue) == cap(ep.user_recv_queue) {
		fmt.Println("warn: processUserData queue is full")
		return
	}
	for {
		data, n := ep.kcp.Recv()
		if n <= 0 {
			break
		}
		select {
		case ep.user_recv_queue <- data:
		default:
			// ERROR
			panic("error: processUserData queue is full")
			return
		}
		if  len(ep.user_recv_queue) == cap(ep.user_recv_queue) {
			fmt.Println("warn: processUserData queue is full")
			return
		}
	}
}

func (ep *KcpEndpoint) Send(data []byte) (err error) {
	ep.output_queue <- data
	return
}

func (ep *KcpEndpoint) SendSYN(data []byte) (err error) {
	ep.kcp.SendSYN(data, len(data))
	return
}

func (ep *KcpEndpoint) Input(data []byte) {
	ep.input_queue <- data
	return
}

func (ep *KcpEndpoint) Close() {
	ep.kcp.SendClose()
}
