package command_test

import (
	"context"
	"testing"

	"github.com/mssvpn/Xray-Lite/app/dispatcher"
	"github.com/mssvpn/Xray-Lite/app/log"
	. "github.com/mssvpn/Xray-Lite/app/log/command"
	"github.com/mssvpn/Xray-Lite/app/proxyman"
	_ "github.com/mssvpn/Xray-Lite/app/proxyman/inbound"
	_ "github.com/mssvpn/Xray-Lite/app/proxyman/outbound"
	"github.com/mssvpn/Xray-Lite/common"
	"github.com/mssvpn/Xray-Lite/common/serial"
	"github.com/mssvpn/Xray-Lite/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
