package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})

	for _, v := range arr {
		m[v] = struct{}{}
	}

	res := make([]string, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}

	fmt.Println(res)
}
