package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/mssvpn/Xray-Lite/app/dispatcher"
	_ "github.com/mssvpn/Xray-Lite/app/proxyman/inbound"
	_ "github.com/mssvpn/Xray-Lite/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/mssvpn/Xray-Lite/app/commander"
	_ "github.com/mssvpn/Xray-Lite/app/log/command"
	_ "github.com/mssvpn/Xray-Lite/app/proxyman/command"
	_ "github.com/mssvpn/Xray-Lite/app/stats/command"

	// Developer preview services
	_ "github.com/mssvpn/Xray-Lite/app/observatory/command"

	// Other optional features.
	_ "github.com/mssvpn/Xray-Lite/app/dns"
	_ "github.com/mssvpn/Xray-Lite/app/dns/fakedns"
	_ "github.com/mssvpn/Xray-Lite/app/log"
	_ "github.com/mssvpn/Xray-Lite/app/metrics"
	_ "github.com/mssvpn/Xray-Lite/app/policy"
	_ "github.com/mssvpn/Xray-Lite/app/reverse"
	_ "github.com/mssvpn/Xray-Lite/app/router"
	_ "github.com/mssvpn/Xray-Lite/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/mssvpn/Xray-Lite/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/mssvpn/Xray-Lite/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/mssvpn/Xray-Lite/proxy/blackhole"
	_ "github.com/mssvpn/Xray-Lite/proxy/dns"
	_ "github.com/mssvpn/Xray-Lite/proxy/dokodemo"
	_ "github.com/mssvpn/Xray-Lite/proxy/freedom"
	_ "github.com/mssvpn/Xray-Lite/proxy/socks"
	_ "github.com/mssvpn/Xray-Lite/proxy/trojan"
	_ "github.com/mssvpn/Xray-Lite/proxy/vless/inbound"
	_ "github.com/mssvpn/Xray-Lite/proxy/vless/outbound"
	_ "github.com/mssvpn/Xray-Lite/proxy/vmess/inbound"
	_ "github.com/mssvpn/Xray-Lite/proxy/vmess/outbound"

	// Transports
	_ "github.com/mssvpn/Xray-Lite/transport/internet/domainsocket"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/grpc"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/http"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/kcp"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/quic"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/reality"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/tcp"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/tls"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/udp"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/websocket"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/xtls"

	// Transport headers
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/http"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/noop"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/srtp"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/tls"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/utp"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/wechat"
	_ "github.com/mssvpn/Xray-Lite/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/mssvpn/Xray-Lite/main/json"
	_ "github.com/mssvpn/Xray-Lite/main/toml"
	_ "github.com/mssvpn/Xray-Lite/main/yaml"

	// Load config from file or http(s)
	_ "github.com/mssvpn/Xray-Lite/main/confloader/external"

	// Commands
	_ "github.com/mssvpn/Xray-Lite/main/commands/all"
)
