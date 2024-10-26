package models

import "time"

type Server struct {
	Name      string    `json:"name"`
	Domain    string    `json:"domain"`
	FQDN      string    `json:"fqdn"`
	Model     string    `json:"model"`
	Role      string    `json:"role"`
	Services  []string  `json:"services"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Networks  []Network     `json:"networks"`
	OpenPorts []PortBinding `json:"open_ports"`

	Type     ServerType `json:"type"`
	Provider Provider   `json:"provider"`

	Location Location `json:"location"`
}

type Provider struct {
	Name string `json:"name"`
}

type ServerType struct {
	Name string `json:"name"`
}

type Network struct {
	Name string   `json:"name"`
	IPv4 []string `json:"ipv4"`
	IPv6 []string `json:"ipv6"`
}

type PortBinding struct {
	Name     string   `json:"name"`
	IPv4     []string `json:"ipv4"`
	IPv6     []string `json:"ipv6"`
	HostPort int      `json:"host_port"`
}

type Location struct {
	DataCenter string `json:"data_center"`
	Rack       string `json:"rack"`
	Row        string `json:"row"`
}
