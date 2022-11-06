package abstractFactory

type DBConnection interface {
	Connection() string
}

type DBCommand interface {
	Command() string
}
type DBRead interface {
	Read() string
}

type DBFactory interface {
	CreateDBConnection() DBConnection
	CreateDBCommand() DBCommand
	CreateDBRead() DBRead
}
