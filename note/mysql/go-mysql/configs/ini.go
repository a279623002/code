package configs

import (
	"fmt"
	"git.chemm.com/backend/lib/config"
	"os"
	"path/filepath"
	"sync"
)

type iniConfig struct {
	mu          sync.RWMutex
	mysqlConfig *MysqlConfig
}

func (ic *iniConfig) reloadConfig() error {
	iniParser := config.IniParser{} // 使用第三方库 go-ini封装 读取ini配置
	workPath, _ := os.Getwd()
	confFileName := filepath.Join(workPath, "conf", "app.conf")
	if err := iniParser.Load(confFileName); err != nil {
		fmt.Printf("try load config file[%s] error[%s]\n", confFileName, err.Error())
		return err
	}

	ic.mu.Lock()

	ic.mysqlConfig.Host = iniParser.GetString("mysql", "host")
	ic.mysqlConfig.Port = int(iniParser.GetInt64("mysql", "port"))
	ic.mysqlConfig.Name = iniParser.GetString("mysql", "name")
	ic.mysqlConfig.UserName = iniParser.GetString("mysql", "username")
	ic.mysqlConfig.Password = iniParser.GetString("mysql", "password")
	ic.mysqlConfig.MaxOpenConns = int(iniParser.GetInt64("mysql", "maxOpenConns"))
	ic.mysqlConfig.MaxIdleConns = int(iniParser.GetInt64("mysql", "maxIdleConns"))

	ic.mu.Unlock()
	return nil
}

var iniconfig *iniConfig

func InitIniConfig() error {
	iniconfig = &iniConfig{
		mu:          sync.RWMutex{},
		mysqlConfig: &MysqlConfig{},
	}
	err := iniconfig.reloadConfig()
	if err != nil {
		return err
	}
	return nil
}

func IniMysqlGet() *MysqlConfig {
	if iniconfig == nil {
		err := InitIniConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	iniconfig.mu.RLock()
	defer iniconfig.mu.RUnlock()
	return iniconfig.mysqlConfig
}
