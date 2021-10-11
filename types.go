package main

type mesh struct {
	Peers map[string]peer `yaml:"peers,omitempty"`
}

type peer struct {
	Description string   `yaml:"description,omitempty"`
	Address     string   `yaml:"address"`
	Entrypoint  string   `yaml:"entrypoint"`
	Port        uint16   `yaml:"port"`
	PrivateKey  string   `yaml:"privatekey"`
	PublicKey   string   `yaml:"publickey"`
	AllowedIPs  []string `yaml:"allowedIPs"`
	// AllowedIPs  []net.IPNet `yaml:"allowedIPs"`
}
