package configs

var Server *server

type server struct {
	Port string
}

func InitServer(cfg *ServerConfig) {
	Server = &server{Port: cfg.Port}
}
