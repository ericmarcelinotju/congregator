package config

type DatabaseList struct {
	Main Database
}

// Database is struct for Database conf
type Database struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Dbname        string
	Schema        string
	Driver        string
	LogLevel      int
	SlowThreshold int
}
