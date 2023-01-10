package strategy

import "fmt"

// 模拟日志记录：文件和数据库存储两种方式
// 抽象日志接口
type Logger interface {
	Info()
	Error()
}

type LogManager struct {
	Logger
}

func NewLogManager(logger Logger) *LogManager {
	return &LogManager{logger}
}

// 文件方式日志
type FileLogger struct {
}

func (f *FileLogger) Info() {
	fmt.Println("文件记录info")
}

func (f *FileLogger) Error() {
	fmt.Println("文件记录Error")
}

// 数据库方式日志
type DbLogger struct {
}

func (d *DbLogger) Info() {
	fmt.Println("数据库记录info")
}

func (d *DbLogger) Error() {
	fmt.Println("数据库记录Error")
}
