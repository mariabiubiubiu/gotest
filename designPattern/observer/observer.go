package observer

import "fmt"

//一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。

type Obserber interface {
	Update(subject *Subject)
}


type Subject struct {
	observers []Obserber
	context string
}

//需要两个操作  第一个 添加/取消订阅 第二个发送更新

func (s *Subject)Attach(o Obserber){
	s.observers = append(s.observers, o)
}

func (s *Subject)notify(){
	for _, o := range s.observers{
		o.Update(s)
	}
}

func (s *Subject)UpdateContext(context string){
	s.context = context
	s.notify()
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}