package main

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {

	for len(this.stack2) > 0 {

		stack2_last_num := len(this.stack2) - 1
		current := this.stack2[stack2_last_num]

		this.stack1 = append(this.stack1, current)
		this.stack2 = this.stack2[:stack2_last_num]
	}
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {

	if len(this.stack1) == 0 {
		stack2_last_num := len(this.stack2) - 1
		first_element := this.stack2[stack2_last_num]
		this.stack2 = this.stack2[:stack2_last_num]

		return first_element
	}

	for len(this.stack1) > 1 {
		stack1_last_num := len(this.stack1) - 1
		current := this.stack1[stack1_last_num]
		this.stack2 = append(this.stack2, current)
		this.stack1 = this.stack1[:stack1_last_num]
	}

	first_element := this.stack1[0]
	this.stack1 = []int{}

	return first_element

}

func (this *MyQueue) Peek() int {

	if len(this.stack1) == 0 {
		stack2_last_num := len(this.stack2) - 1
		first_element := this.stack2[stack2_last_num]

		return first_element
	}

	for len(this.stack1) > 1 {
		stack1_last_num := len(this.stack1) - 1
		current := this.stack1[stack1_last_num]
		this.stack2 = append(this.stack2, current)
		this.stack1 = this.stack1[:stack1_last_num]
	}

	first_element := this.stack1[0]

	return first_element

}

func (this *MyQueue) Empty() bool {
	return (len(this.stack1) == 0) && (len(this.stack2) == 0)
}

func main() {
    //
}
