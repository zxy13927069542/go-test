package main

import "testing"

func TestNilMap(t *testing.T) {
	testNilMap(nil)
}

func testNilMap(m1 map[string]string) {
	println(len(m1))
}
