package gate_listener

import (
	"github.com/tenchlee/roader/endpoint"
	"strconv"
)

type IGateListener interface {
	Listen(port int) error
	Close()
}

type GateListener struct {
	port int
	ip string
	ep_map map[string] *endpoint.IEndpoint
}

func (gl *GateListener)init() {
}

func (gl *GateListener)getEndpoint(id string) (ep *endpoint.IEndpoint) {
	return gl.ep_map[id]
}

func (gl *GateListener)getAddrString() string {
	return ":" + strconv.Itoa(gl.port)
}
