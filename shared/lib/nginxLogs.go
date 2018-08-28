package lib

import (
	"fmt"
)

var nginxSettings interface{}

// initNginx load Nginx configuration
func initNginx() {
	// Récupération de la configuration
	settings := Settings

	nginxSettings = settings.Nginx
}

// GetNginxLogs get Nginx logs
func GetNginxLogs() {
	// Initialisation
	initNginx()

	// Récupération des fichiers de logs

	fmt.Println(nginxSettings)
}
