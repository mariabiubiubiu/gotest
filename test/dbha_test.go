package test

import "testing"

func Test_dbha(t *testing.T)  {
	if e := dbha(); e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
	t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_dbha_1(t *testing.T)  {
	if e := dbha(); e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}