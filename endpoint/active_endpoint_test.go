package endpoint

import (
	"testing"
	"net"
)

func TestRun(t *testing.T) {
	conn, err := net.Dial("udp", "192.168.195.133:8000")
	if err != nil {
		t.Error(err)
		return
	}

}