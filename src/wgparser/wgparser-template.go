package wgparser

var (
	default_fw_rules string = `
	PostUp = iptables -A FORWARD -i <WG-INTERFACE> -d 10.0.0.0/8 -j DROP
	PostUp = iptables -A FORWARD -i <WG-INTERFACE> -d 172.16.0.0/12 -j DROP
	PostUp = iptables -A FORWARD -i <WG-INTERFACE> -d 192.168.0.0/16 -j DROP
	PostUp = iptables -t nat -I POSTROUTING -o <OUTGOING-INTERFACE> -j MASQUERADE
	PreDown = iptables -D FORWARD -i <WG-INTERFACE> -d 10.0.0.0/8 -j DROP
	PreDown = iptables -D FORWARD -i <WG-INTERFACE> -d 172.16.0.0/12 -j DROP
	PreDown = iptables -D FORWARD -i <WG-INTERFACE> -d 192.168.0.0/16 -j DROP
	PreDown = iptables -t nat -D POSTROUTING -o <OUTGOING-INTERFACE> -j MASQUERADE
	`
	interface_template string = `
	[Interface]
	Address = %s
	<FW-RULES>
	ListenPort = %s
	PrivateKey = %s
	`
)
