package main

import (
	"fmt"
	server "mtgbc/server"
)

func main() {
	fmt.Println("Network Address: ", server.MTGNetwork.Address)
	server.StartNode()
}
