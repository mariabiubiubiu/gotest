package vistor

import "fmt"

//访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中。


type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

type CustomerCol struct {
	customers []Customer
}

//添加顾客
func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

//为每个顾客添加访问
func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

//创建一个新用户
func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

//定义接口函数
func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

//私人用户
type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// only for enterprise
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}
