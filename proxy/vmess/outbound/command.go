package outbound

import (
	"time"

	"github.com/mssvpn/Xray-Lite/common"
	"github.com/mssvpn/Xray-Lite/common/net"
	"github.com/mssvpn/Xray-Lite/common/protocol"
	"github.com/mssvpn/Xray-Lite/proxy/vmess"
)

func (h *Handler) handleSwitchAccount(cmd *protocol.CommandSwitchAccount) {
	rawAccount := &vmess.Account{
		Id:      cmd.ID.String(),
		AlterId: uint32(cmd.AlterIds),
		SecuritySettings: &protocol.SecurityConfig{
			Type: protocol.SecurityType_LEGACY,
		},
	}

	account, err := rawAccount.AsAccount()
	common.Must(err)
	user := &protocol.MemoryUser{
		Email:   "",
		Level:   cmd.Level,
		Account: account,
	}
	dest := net.TCPDestination(cmd.Host, cmd.Port)
	until := time.Now().Add(time.Duration(cmd.ValidMin) * time.Minute)
	h.serverList.AddServer(protocol.NewServerSpec(dest, protocol.BeforeTime(until), user))
}

func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch typedCommand := cmd.(type) {
	case *protocol.CommandSwitchAccount:
		if typedCommand.Host == nil {
			typedCommand.Host = dest.Address
		}
		h.handleSwitchAccount(typedCommand)
	default:
	}
}
