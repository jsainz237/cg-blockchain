package main

import (
	"fmt"
	api "mtgbc/api"
	network "mtgbc/network"
)

func main() {
	fmt.Println("Network Address: ", network.MTGNetwork.Address)
	api.Startserver()
}
