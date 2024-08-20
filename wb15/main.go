package main

import (
	"fmt"
)

/*
	=== Задача №15 ===

	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.

	var justString string
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}

*/

/*
	Если взять срез justString = v[:100]
	то эта строка будет ссылаться на всю строку v, то len будет 100, но cap = как у строки v,
	следовательно могут возникнуть утечки памяти при анализе GC, потому что строка justString не отпускает v. А нам в идеале нужны только первые 100 символов.
	Поэтому лучше создать новую строку, чтобы в памяти осталась только нужная часть.
*/

var justString string

func createHugeString(n int) string {
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = byte('S')
	}
	return string(arr)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100]) // создаем новую независимую от v строку
	fmt.Println(len(justString), justString)
	byteString := []byte(justString)
	fmt.Println(cap(byteString)) // удостоверяемся что Cap = 100
}

func main() {
	someFunc()
}
