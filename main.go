package main

import (
	_ "net/http/pprof"
	"github.com/tenchlee/roader/gate_listener"
	"github.com/tenchlee/roader/conn_pool"
)

func main() {
	conn_pool.InitUdpConnPool(":10001")
	tcp_gate := gate_listener.NewTcpGateListener()
	tcp_gate.Listen(7000)
}