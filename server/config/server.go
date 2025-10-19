package config

type ServerList struct {
	Rest   Server
	Gitlab Server
}

type Server struct {
	Host      string
	Port      string
	SecretKey string
	Timeout   int
	Path      string
}
