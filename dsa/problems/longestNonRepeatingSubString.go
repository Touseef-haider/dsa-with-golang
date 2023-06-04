package problems

import (
	"fmt"
	"strings"

	"main.go/utils"
)

func LongestNonRepeatedSubStr(s string) {
	fmt.Println(s)

	var splitted []string = strings.Split(s, "")

	var l, r int

	var windowCount []string = []string{}

	var result []string = []string{}

	for r < len(splitted) {

		var found bool = utils.ElementExistsInArray(windowCount, splitted[r])

		if !found {
			windowCount = append(windowCount, splitted[r])
			if len(windowCount) >= len(result) {
				result = []string{}
				result = append(result, windowCount...)
			}
			r++
		} else {
			fmt.Println("before window", windowCount)

			if len(windowCount) >= len(result) {
				result = []string{}
				result = append(result, windowCount...)
			}
			fmt.Println("result", result)
			data := utils.RemoveIndex(windowCount, 0)

			windowCount = data
			fmt.Println("after window", windowCount)

			l++
		}

	}

	if len(result) == 0 {

		fmt.Println(len(windowCount))

		fmt.Println(windowCount)

	} else {

		fmt.Println(len(result))

		fmt.Println(result)
	}

}
