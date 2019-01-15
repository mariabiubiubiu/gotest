package simpleFactory

//go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。 NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。

type Factory interface {
	Create()string
}

type Car struct {

}

func (self Car) Create() string {
	return "Create Car Done"
}

type Bike struct {

}

func (self Bike) Create()string{
	return "Create Bike Done"
}

func NewFactory(t int)Factory{
	if t == 1{
		return &Car{}
	}else {
		return &Bike{}
	}
}