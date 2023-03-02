package tcp

import (
	"github.com/mssvpn/Xray-Lite/common"
	"github.com/mssvpn/Xray-Lite/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
