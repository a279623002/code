package single

import "sync"

type Single struct {
}

var (
	single *Single
	once   sync.Once
)

func GetHandler() *Single {
	once.Do(func() {
		single = &Single{}
	})
	return single
}
