package abstractFactory

type MysqlConnection struct {
}

func (m *MysqlConnection) Connection() string {
	return "mysql connect"
}

type MysqlCommand struct {
}

func (m *MysqlCommand) Command() string {
	return "mysql command"
}

type MysqlRead struct {
}

func (m *MysqlRead) Read() string {
	return "mysql read"
}

type MysqlFactory struct {
}

func (m *MysqlFactory) CreateDBConnection() DBConnection {
	return &MysqlConnection{}
}
func (m *MysqlFactory) CreateDBCommand() DBCommand {
	return &MysqlCommand{}
}
func (m *MysqlFactory) CreateDBRead() DBRead {
	return &MysqlRead{}
}
