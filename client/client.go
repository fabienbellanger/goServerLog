package client

import (
	"fmt"
	"net/rpc"

	"github.com/fabienbellanger/goServerLog/shared/model"
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
)

// Listen listen RPC server
func Listen(port int) {
	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	client, err := rpc.Dial("tcp", "localhost:"+toolbox.IntToString(port))
	toolbox.CheckError(err, 1)

	// Synchronous call
	log := model.Log{"pos", "onet1", 1, "Message"}

	var name string

	err = client.Call("Log.DisplayProject", log, &name)
	toolbox.CheckError(err, 1)

	fmt.Printf("Projet name: %s\n", name)
}
