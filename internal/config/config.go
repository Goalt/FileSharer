package config

type Config struct {
	DebugLevel int
	Server     Server
}

type Server struct {
	Port int
}
