package etl

import (
	"strings"
)

func Transform(hash map[int][]string) map[string]int {
	m := make(map[string]int)
	for key, value := range hash {
		//fmt.Printf("%T\n", value)
		//fmt.Printf("%T\n ", key)
		for i := 0; i < len(value); i++ {
			a := strings.ToLower(value[i])
			m[a] = key
		}
	}
	//fmt.Print(m)
	return m
}
