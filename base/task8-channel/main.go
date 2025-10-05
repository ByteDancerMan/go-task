package main


func main() { 
	numChan := make(chan int)
	go createNum(numChan) // 在单独的goroutine中运行
	printNum(numChan)
}

func createNum(numChan chan int) {
	for i := 0; i < 10; i++ {
		numChan <- i
	}
	close(numChan) // 发送完毕后关闭channel
}

func printNum(numChan chan int) {
	for num := range numChan { // 使用range遍历channel直到关闭
		println(num)
	}
}
// func main() { 
// 	numChan := make(chan int, 100) // 创建带缓冲的channel
// 	createNum(numChan)
// 	printNum(numChan)
// }

// func createNum(numChan chan int) {
// 	for i := 0; i < 10; i++ {
// 		numChan <- i
// 	}
// 	close(numChan) // 发送完毕后关闭channel
// }

// func printNum(numChan chan int) {
// 	for i := 0; i < 10; i++ {
// 		num := <-numChan
// 		println(num)
// 	}
// }