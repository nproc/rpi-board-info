package native

import (
	"net"

	"github.com/nproc/rpi-board-info/info"
)

func NewNetworkInformationProvider() *NetworkInformationProvider {
	return &NetworkInformationProvider{}
}

type NetworkInformationProvider struct {
}

func (p *NetworkInformationProvider) Networks() ([]info.Network, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	networks := make([]info.Network, len(interfaces))
	for i, inter := range interfaces {
		addrs, err := inter.Addrs()
		if err != nil {
			return nil, err
		}

		ips := make([]string, len(addrs))
		for j, addr := range addrs {
			ips[j] = addr.String()
		}

		networks[i] = info.NewNetwork(
			inter.Name,
			inter.HardwareAddr.String(),
			ips,
		)
	}

	return networks, nil
}
