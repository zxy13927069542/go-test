package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	s1 := `{
		"name": {"first": "Tom", "last": "Anderson"},
		"age":37,
		"children": ["Sara","Alex","Jack"],
		"friends": [
		  {"first": "James", "last": "Murphy"},
		  {"first": "Roger", "last": "Craig"}
		]
	}`

	name := gjson.Get(s1, "name")
	name_last := gjson.Get(s1, "name.last")
	children_1 := gjson.Get(s1, "children.1")
	chil_1 := gjson.Get(s1, "chil*.1")
	chil_ren_1 := gjson.Get(s1, "chil?ren.1")
	children_num := gjson.Get(s1, "children.#")
	children := gjson.Get(s1, "friends.#.first")

	fmt.Println("name:", name)
	fmt.Println("name_last:", name_last)
	fmt.Println("children_1:", children_1)
	fmt.Println("chil_1:", chil_1)
	fmt.Println("chil_ren_1:", chil_ren_1)
	fmt.Println("chilren_num:", children_num)
	fmt.Println("chilren:", children)

}
