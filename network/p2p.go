package network

import (
	_ "mtgbc/env"
	"net"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type Network struct {
	Address        string
	ConnectionPool []string
}

var MTGNetwork Network

func init() {
	ip, err := getOutboundIP()
	if err != nil {
		panic("Could not calculate IP address")
	}

	MTGNetwork = Network{Address: ip}
}

/* Method inspired from https://stackoverflow.com/questions/75268039/how-to-use-golang-to-get-my-wifi-ip-address */
func getOutboundIP() (string, error) {
	port := os.Getenv("PORT")
	ifs, _ := net.Interfaces()

	for _, ifi := range ifs {
		if ifi.Name == "en0" {
			addrs, _ := ifi.Addrs()
			ip := addrs[4].String()
			ip = strings.Split(ip, "/")[0]
			return ip + ":" + port, nil
		}
	}

	return "", errors.New("Could not calculate IP address")
}

func (n *Network) AddConnection(address string) int {
	n.ConnectionPool = append(n.ConnectionPool, address)
	return len(n.ConnectionPool)
}

func (n *Network) RemoveConnection(address string) int {
	for i, connection := range n.ConnectionPool {
		if connection == address {
			n.ConnectionPool = append(n.ConnectionPool[:i], n.ConnectionPool[i+1:]...)
		}
	}

	return len(n.ConnectionPool)
}
