package simpleFactory

type Drive interface {
	Say(string) string
}

func NewDrive(str string) Drive {
	switch str {
	case "cn":
		return &china{}
	case "en":
		return &english{}
	default:
		return nil
	}
}

type china struct {
}

func (*china) Say(name string) string {
	return "cn:" + name
}

type english struct {
}

func (*english) Say(name string) string {
	return "en:" + name
}
