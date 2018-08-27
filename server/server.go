package server

import (
	"net"
	"net/rpc"

	"github.com/fabienbellanger/goServerLog/shared/model"
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
)

// registerLog
func registerLog(server *rpc.Server, log *model.Log) {
	server.RegisterName("Log", log)
}

// Start starts the RPC server
func Start(port int) {
	log := new(model.Log)

	server := rpc.NewServer()
	registerLog(server, log)

	listen, err := net.Listen("tcp", ":"+toolbox.IntToString(port))
	toolbox.CheckError(err, 1)

	server.Accept(listen)
}
