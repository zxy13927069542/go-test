package main

import "fmt"

type Person struct {
	id   int
	name string
}

func (p Person) Print() {
	fmt.Println(p)
}

func (p *Person) setId(id int) {
	p.id = id
}

func main() {
	p := Person{
		id:   1,
		name: "郑晓颖",
	}

	p.Print()
	p.setId(0)

}
