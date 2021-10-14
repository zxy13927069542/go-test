package main

import "fmt"

//结构体继承另一个结构体

type Computer struct {
	*Cpu
	name string
}

//Computer构造函数
func NewComputer(name string, c *Cpu) *Computer {
	return &Computer{
		// Cpu: &Cpu{
		// 	rate: c.rate,
		// },
		Cpu:  c,
		name: name,
	}
}

type Cpu struct {
	rate int
}

//Cpu构造函数
func NewCpu(rate int) *Cpu {
	return &Cpu{
		rate: rate,
	}
}

func (c *Cpu) getRate() int {
	return c.rate
}

func main() {
	c1 := NewCpu(60)
	c2 := NewComputer("DC Pro", c1)
	fmt.Println(c2.getRate())
}
