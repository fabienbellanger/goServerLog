package lib

import (
	"bufio"
	"os"
	"regexp"
	"time"

	"github.com/fabienbellanger/goServerLog/shared/model"
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
func GetNginxLogs() []model.Log {
	// Initialisation
	// --------------
	initSettings()

	// Récupération des logs
	// ---------------------
	logsTmp := getlogsFromFiles()

	return logsTmp
}

// getlogsFromFiles gets logs from files
func getlogsFromFiles() []model.Log {
	yesterdayDate := time.Now().AddDate(0, 0, -1).Format(nginxSettings.DateFormat)
	logs := make([]model.Log, 0)

	// Récupération des fichiers de logs
	// ---------------------------------
	for _, project := range projects {
		// Ouverture du fichier
		// --------------------
		file, err := os.Open(nginxSettings.LogsPaths + "/" + project.Name + nginxSettings.FileSuffix)
		toolbox.CheckError(err, 0)
		defer file.Close()

		// Lecture du fichier
		// ------------------
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			regexPattern := regexp.MustCompile(`(?m)` + yesterdayDate + `.*(Fatal error|Warning|timed out|No database selected).*`)
			line := scanner.Text()
			regexResult := regexPattern.FindAllSubmatch([]byte(line), -1)

			// On ne récupère que les lignes qui nous intéressent
			// --------------------------------------------------
			if len(regexResult) == 1 && len(regexResult[0]) == 2 {
				// On parse la ligne pour récupérer les infos utiles
				// -------------------------------------------------
				regexPattern = regexp.MustCompile(`(\d{4}\/\d{2}\/\d{2})\s(\d{2}:\d{2}:\d{2}).*` +
					`(?:PHP Fatal error:  |PHP Warning:  |timed out |No database selected )` +
					`(?:(.*)(?:, client: )(.*)(?:, server: )(.*)(?:, request: )(.*)(?:, upstream: )(.*)(?:, host: )(.*)|(.*))`)
				regexResult = regexPattern.FindAllSubmatch([]byte(line), -1)

				if len(regexResult) == 1 && len(regexResult[0]) == 10 {
					currentIndex := 0
					linesNumber := len(logs)

					for !(currentIndex == linesNumber || logs[currentIndex].Message == string(regexResult[0][9])) {
						currentIndex++
					}

					if currentIndex == linesNumber {
						// Pas de message trouvé, on ajoute au tableau
						log := model.Log{
							Project:   project.Name,
							Server:    "",
							Number:    1,
							Message:   string(regexResult[0][9]),
							HourStart: string(regexResult[0][2]),
							HourEnd:   string(regexResult[0][2]),
						}

						logs = append(logs, log)
					} else {

						logs[currentIndex].Number++
						logs[currentIndex].HourEnd = string(regexResult[0][2])
					}
				}
			}
		}
	}

	return logs
}
