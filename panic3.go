package main


func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

func main() {
	Try(func() {
		panic("foo")
	}, func(e interface{}) {
		print(e)
	})
}
