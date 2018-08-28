package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/fabienbellanger/goServerLog/shared/toolbox"
)

var nginxSettings Nginx
var projects []Project

// initSettings load configuration
func initSettings() {
	// Récupération de la configuration
	settings := Settings

	nginxSettings = settings.Nginx
	projects = settings.Projects
}

// GetNginxLogs get Nginx logs
func GetNginxLogs() {
	// Initialisation
	// --------------
	initSettings()
	yesterdayDate := time.Now().AddDate(0, 0, -1).Format(nginxSettings.DateFormat)
	log.Println(yesterdayDate)

	// Récupération des fichiers de logs
	// ---------------------------------
	for _, project := range projects {
		log.Println(project.Name)

		// Ouverture du fichier
		// --------------------
		file, err := os.Open(nginxSettings.LogsPaths + "/" + project.Name + nginxSettings.FileSuffix)
		toolbox.CheckError(err, 0)
		defer file.Close()

		// Lecture du fichier
		// ------------------
		lines := make(map[Project][]string)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// GREP Regex : "(\d{4}\/\d{2}\/\d{2})\s(\d{2}:\d{2}:\d{2}).*(?:PHP Fatal error:  |PHP Warning:  |timed out |No database selected )(?:(.*)(?:, client: )(.*)(?:, server: )(.*)(?:, request: )(.*)(?:, upstream: )(.*)(?:, host: )(.*)|(.*))"

			// TODO: regexPattern := regexp.MustCompile(`(?m)` + yesterdayDate + `.*(Fatal error|Warning|timed out|No database selected).*`)
			regexPattern := regexp.MustCompile(`(?m).*(Fatal error|Warning|timed out|No database selected).*`)
			line := scanner.Text()
			regexResult := regexPattern.FindAllSubmatch([]byte(line), -1)

			// On ne récupère que les lignes qui nous intéressent
			if len(regexResult) == 1 && len(regexResult[0]) == 2 {
				lines[project] = append(lines[project], line)

				// On parse la ligne pour récupérer les infos utiles
				// -------------------------------------------------
			}
		}

		log.Println(len(lines[project]))
	}

	log.Println(nginxSettings, projects)
}
