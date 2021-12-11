package main

import (
	"fmt"
	"go-test/ying/float/config"
	"testing"
)

func TestRetainValidNumber(t *testing.T) {
	conf := config.GetConfig()
	nodes := conf.Nodes

	for _, node := range nodes {
		fmt.Println("test case:", node.F, node.N)
		fmt.Println("test result:", RetainValidNumber(node.F, node.N))
	}
}
