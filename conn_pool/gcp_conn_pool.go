package conn_pool

import (
	"net"
	"sync"
	"math"
	"fmt"
	"context"
	"math/rand"
	"time"
	"encoding/binary"
)

type GcpInput interface {
	Input(data []byte)
}

type GcpConn struct {
	Conn net.PacketConn
	lock sync.Mutex
	pool GcpConnPool
	conn_map map[uint32] GcpInput // conv: ep
	conv_sed uint32
	conv_step uint32
	recv_buff_pool sync.Pool
}

type GcpConnPool struct{
	conns []*GcpConn
	addr string
}

func (conn *GcpConn) Close() {
	conn.lock.Lock()
	conn.lock.Unlock()
}

var G_UdpConnPool GcpConnPool

func InitUdpConnPool(addr string) {
	G_UdpConnPool.conns = make([]*GcpConn, 1)
	G_UdpConnPool.addr = addr
	d := net.ListenConfig{
		Control: nil,
	}
	var err error

	for i, _ := range G_UdpConnPool.conns {
		G_UdpConnPool.conns[i] = new(GcpConn)
		rand.Seed(time.Now().UnixNano())
		G_UdpConnPool.conns[i].Conn, err = d.ListenPacket(context.Background(), "udp", addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		G_UdpConnPool.conns[i].conn_map = make(map[uint32] GcpInput)
		G_UdpConnPool.conns[i].conv_sed = rand.Uint32()
		G_UdpConnPool.conns[i].recv_buff_pool.New = func() interface{} {
			return make([]byte, 1500)
		}

		go G_UdpConnPool.conns[i].recv()
	}
}

// todo 改成net.conn 接口
func (pool *GcpConnPool) Get(endpoint GcpInput) (conn *GcpConn, conv uint32)  {
	min := math.MaxInt32
	var min_conn *GcpConn
	for _, conn := range pool.conns {
		conn.lock.Lock()
		if len(conn.conn_map) <= 1 {
			min_conn = conn
			conn.lock.Unlock()
			break
		}
		if len(conn.conn_map) < min {
			min_conn = conn
		}
		conn.lock.Unlock()
	}
	min_conn.lock.Lock()
	conv = min_conn.conv_sed + min_conn.conv_step
	if conv == math.MaxUint32 - 1 {
		min_conn.conv_step = 0
	}
	min_conn.conv_step++
	if min_conn.conn_map[conv] != nil {
		panic("Gcp conn pool get fail")
		return
	}
	min_conn.conn_map[conv] = endpoint
	min_conn.lock.Unlock()
	conn = min_conn
	return
}

func (conn *GcpConn) PutConn(conv uint32) {
	conn.lock.Lock()
	delete(conn.conn_map, conv)
	conn.lock.Unlock()
}


func (conn *GcpConn) recv() {
	for {
		buff := conn.recv_buff_pool.Get().([]byte)
		n, addr, err := conn.Conn.ReadFrom(buff)
		if err != nil {
			break
		}

		// parse kcp data
		var conv uint32
		var cmd uint8
		data := ikcp_decode32u(buff, &conv)
		data = ikcp_decode8u(data, &cmd)
		fmt.Println("recv", conv, cmd, addr, n)
		if conn.conn_map[conv] != nil {
			conn.conn_map[conv].Input(buff[:n])
		} else {
			fmt.Println("recv unkown msg conv", conv)
		}
	}
}

func (conn *GcpConn) PutBuff(data []byte) {
	data = data[:1500]
	conn.recv_buff_pool.Put(data)
}

/* decode 32 bits unsigned int (lsb) */
func ikcp_decode32u(p []byte, l *uint32) []byte {
	*l = binary.LittleEndian.Uint32(p)
	return p[4:]
}

/* decode 8 bits unsigned int */
func ikcp_decode8u(p []byte, c *byte) []byte {
	*c = p[0]
	return p[1:]
}
