package model

// Server type
type Server string

// Project type
type Project string

// Log type
type Log struct {
	Project Project
	Server  Server
	Number  int
	Message string
	// HourStart time.Time
	// HourEnd   time.Time
}

// DisplayProject display log project name
func (l *Log) DisplayProject(log *Log, name *Project) error {
	*name = log.Project

	return nil
}
