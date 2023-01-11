package proxy

type RelSubject struct {
}

func (r *RelSubject) Do() string {
	return "hello"
}

type Proxy struct {
	rel   *RelSubject
	money int
}

func (p *Proxy) Do() string {
	if p.money > 0 {
		return p.rel.Do()
	}
	return "not enough"
}
