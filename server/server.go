package server

import (
	"net"
	"net/rpc"

	"github.com/fabienbellanger/goServerLog/shared/model"
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
)

func registerLog(server *rpc.Server, log *model.Log) {
	// registers Log interface by name of `Log`.
	// If you want this name to be same as the type name, you
	// can use server.Register instead.
	server.RegisterName("Log", log)
}

// Start starts the RPC server
func Start(port int) {
	log := new(model.Log)

	// Register a new rpc server (In most cases, you will use default server only)
	// And register struct we created above by name "Log"
	// The wrapper method here ensures that only structs which implement Log interface
	// are allowed to register themselves.
	server := rpc.NewServer()
	registerLog(server, log)

	// Listen for incoming tcp packets on specified port.
	listen, err := net.Listen("tcp", ":"+toolbox.IntToString(port))
	toolbox.CheckError(err, 1)

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	server.Accept(listen)
}
