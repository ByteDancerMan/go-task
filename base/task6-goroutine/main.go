package main

import "time"

func main() {
	go printEvenNum()
	go printOddNum()
	time.Sleep(1 * time.Second)
}

func printOddNum() {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			println(i)
		}
	}
}

func printEvenNum() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}
}