package client

import (
	"fmt"
	"net/rpc"

	"github.com/fabienbellanger/goServerLog/shared/lib"
	"github.com/fabienbellanger/goServerLog/shared/model"
	"github.com/fabienbellanger/goServerLog/shared/toolbox"
)

// Listen listen RPC server
func Listen(port int) {
	// Tries to connect to localhost:1234 (The port on which rpc server is listening)
	client, err := rpc.Dial("tcp", "localhost:"+toolbox.IntToString(port))
	toolbox.CheckError(err, 1)

	log := model.Log{
		Project:   "pos",
		Server:    "onet1",
		Number:    1,
		Message:   "Message",
		HourStart: "12:04:12",
		HourEnd:   "22:56:54"}

	var logStr string

	err = client.Call("Log.DisplayLog", log, &logStr)
	toolbox.CheckError(err, 1)

	fmt.Printf("Log: %s\n", logStr)
}

// GetLogs get logs
func GetLogs() {
	// Nginx logs
	// ----------
	lib.GetNginxLogs()
}
