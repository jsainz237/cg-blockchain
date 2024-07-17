package api

import (
	"fmt"
	"log"
	network "mtgbc/network"
	"net/http"
	"net/rpc"
	"slices"
)

type NetworkHandlers struct{}

func registerNodeToConnection(connection string, nodeUrl string) (string, error) {
	var clientReply string
	client, err := rpc.DialHTTP("tcp", connection)

	if err != nil {
		return "", err
	}

	err = client.Call("Network.Register", &ConnectArgs{NodeUrl: nodeUrl}, clientReply)

	if err != nil {
		return "", err
	}

	return clientReply, nil
}

func registerExistingConnectionsToNode(nodeUrl string, allUrls []string) (string, error) {
	log.Printf("[Info] All URLs: %v", allUrls)

	// Register all other nodes on the new node
	client, err := rpc.DialHTTP("tcp", nodeUrl)
	if err != nil {
		return "", err
	}

	var clientReply string
	err = client.Call("Network.RegisterBulk", &RegisterBulkArgs{NodeUrls: allUrls}, clientReply)
	if err != nil {
		return "", err
	}

	return clientReply, nil
}

type ConnectArgs struct {
	NodeUrl string `json:"nodeUrl"`
}

func (nh *NetworkHandlers) Connect(r *http.Request, args *ConnectArgs, reply *string) error {
	nodeUrl := args.NodeUrl

	if slices.Contains(network.MTGNetwork.ConnectionPool, nodeUrl) {
		return fmt.Errorf("node already connected")
	}

	// Register all existing connections on the new node
	allUrls := append(network.MTGNetwork.ConnectionPool[:], network.MTGNetwork.Address)
	_, err := registerExistingConnectionsToNode(nodeUrl, allUrls)

	if err != nil {
		return fmt.Errorf("[Error] Could not register existing nodes on %s: %s", nodeUrl, err.Error())
	}

	// Register node on all other nodes in the network
	for _, connection := range network.MTGNetwork.ConnectionPool {
		_, err := registerNodeToConnection(connection, nodeUrl)
		if err != nil {
			log.Printf("[Error] Could not register node on %s: %s", connection, err.Error())
			continue
		}
	}

	// Add node to the current network
	network.MTGNetwork.AddConnection(nodeUrl)
	return nil
}

func (nh *NetworkHandlers) Register(r *http.Request, args *ConnectArgs, reply *string) error {
	nodeUrl := args.NodeUrl

	log.Printf("[Info] Registering node %s", nodeUrl)
	network.MTGNetwork.AddConnection(nodeUrl)
	*reply = "Node registered on network"
	return nil
}

type RegisterBulkArgs struct {
	NodeUrls []string `json:"nodeUrls"`
}

func (nh *NetworkHandlers) RegisterBulk(r *http.Request, args *RegisterBulkArgs, reply *string) error {
	nodeUrls := args.NodeUrls

	for _, nodeUrl := range nodeUrls {
		log.Printf("[Info] Registering node %s via bulk", nodeUrl)
		network.MTGNetwork.AddConnection(nodeUrl)
	}

	*reply = "All nodes registered on network"
	return nil
}

func (nh *NetworkHandlers) GetConnection(r *http.Request, args *struct{}, reply *[]string) error {
	if len(network.MTGNetwork.ConnectionPool) == 0 {
		*reply = []string{}
		return nil
	}

	*reply = network.MTGNetwork.ConnectionPool
	return nil
}
