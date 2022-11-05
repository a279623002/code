package shiro

type Shiro struct {
	localAdapter
}

type localAdapter = Adapter

func New() *Shiro {
	// default adapter
	s := &Shiro{}
	return s
}

func (s *Shiro) SetAdapter(adapter Adapter) {
	s.localAdapter = adapter
}
