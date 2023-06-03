package problems

import (
	"fmt"
	"strconv"

	linkedlist "main.go/dsa/linkedList"
)

func AddTwoNumbersProb() {
	fmt.Println("Add two numbers")

	var l1 linkedlist.LinkedList = linkedlist.LinkedList{}

	l1.AddNode(2)
	l1.AddNode(4)
	l1.AddNode(3)

	var l2 linkedlist.LinkedList = linkedlist.LinkedList{}

	l2.AddNode(5)
	l2.AddNode(6)
	l2.AddNode(4)

	var l1Head *linkedlist.Node = l1.Reverse()
	var l2Head *linkedlist.Node = l2.Reverse()

	temp1 := l1Head
	temp2 := l2Head

	var res1 string
	var res2 string

	for temp1 != nil {
		res1 += strconv.Itoa(temp1.GetData())
		temp1 = temp1.ReturnNext()
	}
	for temp2 != nil {
		res2 += strconv.Itoa(temp2.GetData())
		temp2 = temp2.ReturnNext()
	}

	fmt.Println(res1)
	fmt.Println(res2)

	num1, _ := strconv.Atoi(res1)
	num2, _ := strconv.Atoi(res2)

	var result int = num1 + num2

	var resultant linkedlist.LinkedList = linkedlist.LinkedList{}

	temp := result

	fmt.Println(temp)

	for temp != 0 {
		resultant.AddNode(temp % 10)
		temp = temp / 10

	}

	resultant.Traverse()

}
