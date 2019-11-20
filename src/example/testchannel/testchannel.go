package main

import "fmt"

// createChannel 接受一个int类型的参数，返回一个int类型的通道
func createChannel(startInt int, chanName string) chan int {
	next := make(chan int)
	// 建立新的goruntime
	go func(i int, name string) {
		for {
			next <- i
			fmt.Printf("ChannelName: %8s is now adding number %d\n", name, i)
			i++
		}
	}(startInt, chanName)
	// 由于chan是引用类型，所以返回以后在子goruntime还会继续运行
	return next
}

func main() {
	counterA := createChannel(8, "ChannelA")
	counterB := createChannel(108, "ChannelB")
	for i := 0; i < 20; i++ {
		a := <-counterA
		fmt.Printf("(A->%d, B->%d)\n", a, <-counterB)
	}
}
