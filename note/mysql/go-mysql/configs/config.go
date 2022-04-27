package configs

type MysqlConfig struct {
	Host         string
	Port         int
	Name         string
	UserName     string
	Password     string
	MaxOpenConns int // 设置的最大连接数 当执行完sql，连接转移到rows对象上，如果rows不关闭，这条连接不会被放回池里，其他并发获取不到连接会被阻塞住
	MaxIdleConns int // 设置的执行完闲置的连接
}
