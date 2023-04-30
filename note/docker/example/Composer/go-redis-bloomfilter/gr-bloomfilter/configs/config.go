package configs

type SessionConfig struct {
	Host       string
	Port       string
	Select     int
	Authstring string
	BFFunc     string // 指定hash算法
	BFBucket   string // bucket
}

type ServerConfig struct {
	Port string
}
