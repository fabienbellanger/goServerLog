package lib

import (
	"encoding/json"
	"os"

	"github.com/fabienbellanger/goServerLog/shared/toolbox"
)

// Config type
type Config struct {
	Nginx struct {
		LogsPaths  string
		FileSuffix string
	}
}

// Settings global variable
var Settings Config

// Init : Lecture du fichier de configuration
func Init() {
	// Lecture du fichier de configuration
	// -----------------------------------
	file, _ := os.Open("settings.json")
	defer file.Close()

	// Décodage du JSON
	// ----------------
	decoder := json.NewDecoder(file)
	settings := Config{}
	err := decoder.Decode(&settings)
	toolbox.CheckError(err, 1)

	Settings = settings
}
