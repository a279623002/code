package configs

import (
	"github.com/go-ini/ini"
	"sync"
)

type iniConfig struct {
	mu            sync.RWMutex
	sessionConfig *SessionConfig
	serverConfig  *ServerConfig
}

func (ic *iniConfig) reloadConfig() error {
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		return err
	}
	ic.mu.Lock()
	defer ic.mu.Unlock()

	ic.sessionConfig.Host = cfg.Section("redis").Key("host").String()
	ic.sessionConfig.Port = cfg.Section("redis").Key("port").String()
	ic.sessionConfig.Select, err = cfg.Section("redis").Key("select").Int()
	if err != nil {
		return err
	}
	ic.sessionConfig.Authstring = cfg.Section("redis").Key("auth_string").String()
	ic.sessionConfig.BFFunc = cfg.Section("redis").Key("bf_func").String()
	ic.sessionConfig.BFBucket = cfg.Section("redis").Key("bf_bucket").String()

	ic.serverConfig.Port = cfg.Section("server").Key("port").String()

	return err
}

var conf *iniConfig

func InitIniConfig() error {
	conf = &iniConfig{
		mu:            sync.RWMutex{},
		sessionConfig: &SessionConfig{},
		serverConfig:  &ServerConfig{},
	}
	err := conf.reloadConfig()
	if err != nil {
		return err
	}
	return nil
}

func IniSessionGet() (cfg *SessionConfig, err error) {
	if conf == nil {
		err := InitIniConfig()
		if err != nil {
			return nil, err
		}
	}
	conf.mu.RLock()
	defer conf.mu.RUnlock()
	return conf.sessionConfig, nil
}

func IniServerGet() (cfg *ServerConfig, err error) {
	if conf == nil {
		err := InitIniConfig()
		if err != nil {
			return nil, err
		}
	}
	conf.mu.RLock()
	defer conf.mu.RUnlock()
	return conf.serverConfig, err
}
