package api

import (
	_ "mtgbc/env"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func Startserver() {
	port := os.Getenv("PORT")

	_, router := initRoutes()
	http.ListenAndServe(":"+port, router)
}

func initRoutes() (*rpc.Server, *mux.Router) {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	blockchainHandlers := new(BlockchainHandlers)
	server.RegisterService(blockchainHandlers, "Blockchain")

	router := mux.NewRouter()
	router.Handle("/rpc", server)

	return server, router
}
