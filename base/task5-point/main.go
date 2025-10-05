package main


func main() {
	num := 5
	add(&num)
	println(num)

	println("-------------------------")

	nums := []int{1, 2, 3, 4, 5}
	mul2(&nums)
	for num := range nums {
		println(nums[num])
	}
}

// add 函数接受一个整数指针作为参数，并返回 num 加上 10 的结果
func add(num *int) {
	*num = *num + 10
}

// mul2 函数接受一个整数切片指针作为参数，并返回切片的每个元素乘以 2 的结果
func mul2(num *[]int) {
	for index := range *num {
		(*num)[index] = (*num)[index] * 2
	}
}