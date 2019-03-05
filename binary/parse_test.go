package binary

import (
	"testing"
	"math"
	"fmt"
)

func TestHtonl(t *testing.T) {
	in := uint32(1341324123)
	out := Htonl(in)
	if out != 1543238223 {
		t.Error("error 123456")
	}

	in = uint32(0)
	out = Htonl(in)
	if out != 0 {
		t.Error("error 0")
	}
}

func TestHtons(t *testing.T) {
	in := uint16(80)
	out := Htons(in)
	if out != 20480 {
		t.Error("error 80")
	}

	in = uint16(0)
	out = Htons(in)
	if out != 0 {
		t.Error("error 0")
	}
}

func TestIPv4IPtoUint(t *testing.T) {
	result := IPv4IPtoUint("172.16.1.5")
	if result != 83955884 {
		t.Error("error 0")
	}
	result = IPv4IPtoUint("1172.16.1.5")
	if result != math.MaxUint32 {
		t.Error("error 1")
	}
}

func TestInetNtoa(t *testing.T) {
	result := InetNtoa(16777343)
	fmt.Println(result)
}

func TestAA(t *testing.T) {
	buff := make([]byte, 1500)
	buff = buff[:10]
	buff = buff[:1500]
	buff[1499] = 'a'
}
