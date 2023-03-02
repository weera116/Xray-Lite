package mux_test

import (
	"testing"

	"github.com/mssvpn/Xray-Lite/common"
	"github.com/mssvpn/Xray-Lite/common/buf"
	"github.com/mssvpn/Xray-Lite/common/mux"
	"github.com/mssvpn/Xray-Lite/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.example.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
