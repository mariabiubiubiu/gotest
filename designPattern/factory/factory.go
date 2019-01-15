package factory

//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result()int
}

//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

//OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

//SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

//SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//PlusOperator Operator 的实际加法实现 extend
type PlusOperator struct {
	*OperatorBase
}

//PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

type SubOperator struct {
	*OperatorBase
}

func (o SubOperator) Result()int{
	return o.a - o.b
}

type SubOperatorFactory struct {
}

func (SubOperatorFactory) Create()Operator  {
	return &SubOperator{
		OperatorBase:&OperatorBase{},
	}
}

