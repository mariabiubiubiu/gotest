package facade

import "fmt"

//facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。

type aModuleImpl struct{}

type bModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

type BModuleAPI interface {
	TestB() string
}

type AModuleAPI interface {
	TestA() string
}

//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

//API is facade interface of facade package
type API interface {
	Test() string
}

//facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (api *apiImpl) Test() string {
	aRet := api.a.TestA()
	bRet := api.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}