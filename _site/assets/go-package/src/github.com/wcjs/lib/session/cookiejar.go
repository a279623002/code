package session

import (
	"net/http/cookiejar"
	"sync"
)

type CookiejarManger struct {
	Mu   sync.RWMutex
	Data map[string]*cookiejar.Jar
}

var cookiejarManger = &CookiejarManger{
	Data: make(map[string]*cookiejar.Jar, 0),
}

func NewCookiejar() *CookiejarManger {
	return cookiejarManger
}

func (cj *CookiejarManger) Create() (jar *cookiejar.Jar, sid string) {
	jar, _ = cookiejar.New(nil)
	sid = GenerateSID()
	cj.Set(sid, jar)
	return jar, sid
}

func (cj *CookiejarManger) Get(sid string) *cookiejar.Jar {
	cj.Mu.RLock()
	defer cj.Mu.RUnlock()

	jar, ok := cj.Data[sid]
	if ok {
		return jar
	} else {
		return nil
	}
}

func (cj *CookiejarManger) Set(sid string, jar *cookiejar.Jar) {
	cj.Mu.Lock()
	cj.Data[sid] = jar
	cj.Mu.Unlock()
}

func (cj *CookiejarManger) Gc() {
	//@todo
}