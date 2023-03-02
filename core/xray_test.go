package core_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/mssvpn/Xray-Lite/app/dispatcher"
	"github.com/mssvpn/Xray-Lite/app/proxyman"
	"github.com/mssvpn/Xray-Lite/common"
	"github.com/mssvpn/Xray-Lite/common/net"
	"github.com/mssvpn/Xray-Lite/common/protocol"
	"github.com/mssvpn/Xray-Lite/common/serial"
	"github.com/mssvpn/Xray-Lite/common/uuid"
	. "github.com/mssvpn/Xray-Lite/core"
	"github.com/mssvpn/Xray-Lite/features/dns"
	"github.com/mssvpn/Xray-Lite/features/dns/localdns"
	_ "github.com/mssvpn/Xray-Lite/main/distro/all"
	"github.com/mssvpn/Xray-Lite/proxy/dokodemo"
	"github.com/mssvpn/Xray-Lite/proxy/vmess"
	"github.com/mssvpn/Xray-Lite/proxy/vmess/outbound"
	"github.com/mssvpn/Xray-Lite/testing/servers/tcp"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
