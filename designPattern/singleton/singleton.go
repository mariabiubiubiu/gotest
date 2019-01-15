package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
//对于从全局的角度只需要运行一次的代码，比如全局初化操始作，go语言提供了一个Once类型来保证全局的唯一性操作。
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}