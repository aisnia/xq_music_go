package mymiddleware

import (
	"fmt"
	"testing"
	"time"
)

func TestWorker_GetId(t *testing.T) {
	// 生成节点实例
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}
	for{
		time.Sleep(time.Millisecond * 10)
		fmt.Println(node.GetId())
	}
}
