package taggedimpl

import (
	"context"

	"github.com/mssvpn/Xray-Lite/common/net"
	"github.com/mssvpn/Xray-Lite/common/net/cnc"
	"github.com/mssvpn/Xray-Lite/common/session"
	"github.com/mssvpn/Xray-Lite/core"
	"github.com/mssvpn/Xray-Lite/features/routing"
	"github.com/mssvpn/Xray-Lite/transport/internet/tagged"
)

func DialTaggedOutbound(ctx context.Context, dest net.Destination, tag string) (net.Conn, error) {
	var dispatcher routing.Dispatcher
	if core.FromContext(ctx) == nil {
		return nil, newError("Instance context variable is not in context, dial denied. ")
	}
	if err := core.RequireFeatures(ctx, func(dispatcherInstance routing.Dispatcher) {
		dispatcher = dispatcherInstance
	}); err != nil {
		return nil, newError("Required Feature dispatcher not resolved").Base(err)
	}

	content := new(session.Content)
	content.SkipDNSResolve = true

	ctx = session.ContextWithContent(ctx, content)
	ctx = session.SetForcedOutboundTagToContext(ctx, tag)

	r, err := dispatcher.Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt cnc.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = cnc.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = cnc.ConnectionOutputMultiUDP(r.Reader)
	}
	return cnc.NewConnection(cnc.ConnectionInputMulti(r.Writer), readerOpt), nil
}

func init() {
	tagged.Dialer = DialTaggedOutbound
}
