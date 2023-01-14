package observer

import "container/list"

// 新闻主题
type NewsSubject struct {
	tiitle string
	l      *list.List
}

func (s *NewsSubject) Add(o Observer) {
	s.l.PushBack(o)
}

func (s *NewsSubject) Send(str string) {
	for i := s.l.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(s.tiitle + "send: " + str)
	}
}

// 热点主题
type HotSubject struct {
	tiitle string
	l      *list.List
}

func (h *HotSubject) Add(o Observer) {
	h.l.PushBack(o)
}

func (h *HotSubject) Send(str string) {
	for i := h.l.Front(); i != nil; i = i.Next() {
		(i.Value).(Observer).Receive(h.tiitle + "send: " + str)
	}
}
