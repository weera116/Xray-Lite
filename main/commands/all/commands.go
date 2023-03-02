package all

import (
	"github.com/mssvpn/Xray-Lite/main/commands/all/api"
	"github.com/mssvpn/Xray-Lite/main/commands/all/tls"
	"github.com/mssvpn/Xray-Lite/main/commands/base"
)

// go:generate go run github.com/mssvpn/Xray-Lite/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
