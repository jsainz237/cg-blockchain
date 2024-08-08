package api

import (
	"fmt"
	"log"
	network "mtgbc/network"
	"net/http"
	"slices"
)

type NetworkHandlers struct{}

func (nh *NetworkHandlers) Ping(r *http.Request, args *struct{}, reply *string) error {
	*reply = "OK"
	return nil
}

func (nh *NetworkHandlers) ConnectPeer(r *http.Request, args *string, reply *string) error {
	nodeUrl := *args

	if slices.Contains(network.MTGNetwork.ConnectionPool, nodeUrl) {
		return fmt.Errorf("node already connected")
	}

	// Register all existing connections on the new node
	allUrls := append(network.MTGNetwork.ConnectionPool[:], network.MTGNetwork.Address)
	_, _, err := CallRPC(nodeUrl, "Network.RegisterBulk", &allUrls)
	if err != nil {
		return fmt.Errorf("[Error] Could not register existing nodes on %s: %s", nodeUrl, err.Error())
	}

	// Register node on all other nodes in the network
	for _, connection := range network.MTGNetwork.ConnectionPool {
		_, _, err := CallRPC(connection, "Network.Register", nodeUrl)
		if err != nil {
			log.Printf("[Error] Could not register node on %s: %s", connection, err.Error())
			continue
		}
	}

	// Add node to the current network
	network.MTGNetwork.AddConnection(nodeUrl)
	*reply = "Node connected to network"
	return nil
}

func (nh *NetworkHandlers) Register(r *http.Request, args *string, reply *string) error {
	nodeUrl := *args

	log.Printf("[Info] Registering node %s", nodeUrl)
	network.MTGNetwork.AddConnection(nodeUrl)
	*reply = "Node registered on network"
	return nil
}

func (nh *NetworkHandlers) RegisterBulk(r *http.Request, args *[]string, reply *string) error {
	nodeUrls := *args

	for _, nodeUrl := range nodeUrls {
		log.Printf("[Info] Registering node %s via bulk", nodeUrl)
		network.MTGNetwork.AddConnection(nodeUrl)
	}

	*reply = "All nodes registered on network"
	return nil
}

func (nh *NetworkHandlers) GetConnections(r *http.Request, args *struct{}, reply *[]string) error {
	if len(network.MTGNetwork.ConnectionPool) == 0 {
		*reply = []string{}
		return nil
	}

	*reply = network.MTGNetwork.ConnectionPool
	return nil
}
