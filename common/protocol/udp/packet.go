package udp

import (
	"github.com/mssvpn/Xray-Lite/common/buf"
	"github.com/mssvpn/Xray-Lite/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
