package observer

import (
	"container/list"
	"testing"
)

func TestNewsSubject_Send(t *testing.T) {
	a := &aObserver{name: "zzq"}
	b := &bObserver{name: "shiro"}

	news := NewsSubject{
		tiitle: "新闻",
		l:      list.New(),
	}
	news.Add(a)
	news.Add(b)

	news.Send("天气预报，在前方3公里下起暴雨")

	hot := HotSubject{
		tiitle: "热点",
		l:      list.New(),
	}
	hot.Add(a)

	hot.Send("b没订阅收不到")
}
