package api

import (
	"bytes"
	"encoding/json"
	"log"
	network "mtgbc/network"
	"net/http"
	"slices"

	"github.com/labstack/echo/v4"
)

type ConnectHandlerRequest struct {
	NodeUrl string `json:"nodeUrl"`
}

var connectHandler = func(c echo.Context) error {
	connectHandlerRequest := new(ConnectHandlerRequest)
	if err := c.Bind(connectHandlerRequest); err != nil {
		log.Printf("[Error] Could not bind nodeUrl: %s", err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	nodeUrl := connectHandlerRequest.NodeUrl

	if slices.Contains(network.MTGNetwork.ConnectionPool, nodeUrl) {
		return c.JSON(http.StatusBadRequest, "Node already connected")
	}

	reqBody, _ := json.Marshal(map[string]string{
		"nodeUrl": nodeUrl,
	})

	// Register node on all other nodes in the network
	for _, connection := range network.MTGNetwork.ConnectionPool {
		_, err := http.Post(connection+"/register", "application/json", bytes.NewBuffer(reqBody))
		if err == nil {
			log.Fatalf("Could not register node on %s", connection)
		}
	}

	// Add node to the current network
	network.MTGNetwork.AddConnection(nodeUrl)
	return c.JSON(http.StatusCreated, "Node connected to network")
}

var registerHandler = func(c echo.Context) error {
	connectHandlerRequest := new(ConnectHandlerRequest)
	if err := c.Bind(connectHandlerRequest); err != nil {
		log.Printf("[Error] Could not bind nodeUrl: %s", err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	nodeUrl := connectHandlerRequest.NodeUrl

	network.MTGNetwork.AddConnection(nodeUrl)
	return c.JSON(http.StatusCreated, "Node registered on network")
}

var getConnectionsHandler = func(c echo.Context) error {
	if len(network.MTGNetwork.ConnectionPool) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	return c.JSON(http.StatusOK, network.MTGNetwork.ConnectionPool)
}
