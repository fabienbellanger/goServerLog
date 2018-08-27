package model

import "github.com/fabienbellanger/goServerLog/shared/toolbox"

// Log type
type Log struct {
	Project   string
	Server    string
	Number    int
	Message   string
	HourStart string
	HourEnd   string
}

// DisplayLog display log
func (l *Log) DisplayLog(log *Log, str *string) error {
	*str = toolbox.IntToString(log.Number) + " x "
	*str += "[" + log.HourStart + " - " + log.HourEnd + "] "
	*str += log.Project + " : " + log.Message

	return nil
}

// SendLogs send logs to server
func (l *Log) SendLogs(log *[]Log, success *bool) error {
	// TODO: Enregistrement des logs et envoi par mail

	*success = true

	return nil
}
