package client

import (
	"log"

	"github.com/fabienbellanger/goServerLog/shared/lib"
)

// GetLogs get logs
func GetLogs() {
	// Nginx logs
	// ----------
	log.Println(lib.GetNginxLogs())
}
