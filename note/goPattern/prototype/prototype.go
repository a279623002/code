package prototype

type Client interface {
	Clone() Client
}

type PrototypeManage struct {
	ClientArr map[string]Client
}

func NewPrototypeManage() *PrototypeManage {
	return &PrototypeManage{
		make(map[string]Client),
	}
}
