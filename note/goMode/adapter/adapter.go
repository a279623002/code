package adapter

type Adapter interface {
	Set(string)
	Get() string
}
