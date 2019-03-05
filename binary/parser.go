package binary

import (
	"net"
	"math"
	"fmt"
)

func Htonl(in uint32) (out uint32) {
	out = uint32(byte(in >> 24)) | (uint32(byte(in >> 16)) << 8)| (uint32(byte(in >> 8)) << 16) | (uint32(byte(in)) << 24)
	return
}

func Htons(in uint16) (out uint16) {
	out = (uint16(byte(in >> 8))) | (uint16(byte(in)) << 8)
	return
}

func IPv4IPtoUint(ip string) uint32 {
	mip := net.ParseIP(ip).To4()
	if mip == nil {
		return math.MaxUint32
	}
	var sum uint32
	sum += uint32(mip[3]) << 24
	sum += uint32(mip[2]) << 16
	sum += uint32(mip[1]) << 8
	sum += uint32(mip[0])

	return sum
}

func InetNtoa(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip), byte(ip >> 8), byte(ip >> 16), byte(ip >> 24))
}