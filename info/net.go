package info

type Network interface {
	Name() string
	HardwareAddress() string
	IPs() []string
}

type NetworkInformationProvider interface {
	Networks() ([]Network, error)
}

func NewNetwork(name, hardwareAddress string, ips []string) Network {
	return &network{
		name:            name,
		hardwareAddress: hardwareAddress,
		ips:             ips,
	}
}

type network struct {
	name            string
	hardwareAddress string
	ips             []string
}

func (n *network) Name() string {
	return n.name
}

func (n *network) HardwareAddress() string {
	return n.hardwareAddress
}

func (n *network) IPs() []string {
	return n.ips
}
