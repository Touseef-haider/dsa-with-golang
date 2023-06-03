package problems

import (
	"fmt"
	"strings"
)

func LongestNonRepeatedSubStr(s *string) {

	var res string

	var isRepeated bool = false

	var hm map[string]string = map[string]string{}

	var splitted = strings.Split(*s, "")

	fmt.Println(res, hm, splitted, len(splitted))

	if len(*s) == 1 {
		return
	}

	for i := 0; i < len(splitted); i++ {
		if hm[splitted[i]] != splitted[i] {
			hm[splitted[i]] = splitted[i]
		} else {
			isRepeated = true

			if len(res) < len(hm) {
				res = ""
				for value, _ := range hm {
					res += value
				}
			}

			hm = map[string]string{}
			hm[splitted[i]] = splitted[i]

		}
	}

	if !isRepeated {
		for value, _ := range hm {
			res += value
		}

		return
	}

	fmt.Println(res)

}
