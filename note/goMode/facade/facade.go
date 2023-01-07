package facade

import "fmt"

type API interface {
	Test() string
}

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct {
}

func (a *aModuleImpl) TestA() string {
	return "shiro"
}

func NewAModuleApi() AModuleAPI {
	return &aModuleImpl{}
}

type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct {
}

func (b *bModuleImpl) TestB() string {
	return "zzq"
}

func NewBModuleApi() BModuleAPI {
	return &bModuleImpl{}
}

type APICall struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *APICall) Test() string {
	return fmt.Sprintf("%s and %s", a.a.TestA(), a.b.TestB())
}
