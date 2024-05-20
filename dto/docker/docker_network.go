package docker

import (
	"docker-site/dto"
	"time"
)

type DockerNetwork struct {
	Name       string    `json:"Name"`
	Id         string    `json:"Id"`
	Created    time.Time `json:"Created"`
	Scope      string    `json:"Scope"`
	Driver     string    `json:"Driver"`
	EnableIPv6 bool      `json:"EnableIPv6"`
	Internal   bool      `json:"Internal"`
	Attachable bool      `json:"Attachable"`
	Ingress    bool      `json:"Ingress"`
}

func (c *DockerNetwork) CountElement(network *dto.NetworkResume) {
	switch c.Driver {
	case "host":
		network.NbHost += 1
		break
	case "bridge":
		network.NbBridge += 1
		break
	case "none":
		network.NbNone += 1
		break
	case "overlay":
		network.NbOverlay += 1
		break
	case "macvlan":
		network.NbMacVlan += 1
		break
	case "ipvlan":
		network.NbIpVlan += 1
		break
	default:
		return
	}
}
