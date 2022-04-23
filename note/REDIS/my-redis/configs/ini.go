package configs

import (
	"fmt"
	"git.chemm.com/backend/lib/config"
	"os"
	"path/filepath"
	"sync"
)

type iniConfig struct {
	mu            sync.RWMutex
	sessionConfig *SessionConfig
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

	ic.sessionConfig.Host = iniParser.GetString("session", "host")
	ic.sessionConfig.Select = int(iniParser.GetInt64("session", "select"))
	ic.sessionConfig.Authstring = iniParser.GetString("session", "auth_string")

	ic.mu.Unlock()
	return nil
}

var iniconfig *iniConfig

func InitIniConfig() error {
	iniconfig = &iniConfig{
		mu:            sync.RWMutex{},
		sessionConfig: &SessionConfig{},
	}
	err := iniconfig.reloadConfig()
	if err != nil {
		return err
	}
	return nil
}

func IniSessionGet() *SessionConfig {
	if iniconfig == nil {
		err := InitIniConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	iniconfig.mu.RLock()
	defer iniconfig.mu.RUnlock()
	return iniconfig.sessionConfig
}
