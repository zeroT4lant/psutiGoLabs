package main

import (
	"fmt"
	"psutiGoLabs/laba3/mathutils"
	"psutiGoLabs/laba3/stringutils"
	"unicode/utf8"
)

func main() {
	//region #1-2 Exercise
	fmt.Println(mathutils.Factorial(5))

	var someNum int
	fmt.Scan(&someNum)
	fmt.Println(mathutils.Factorial(someNum))

	//region #3 Exercise
	stringutils.ReverseString("abobus")

	//region #4 Exercise
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	//region #5 Exercise
	massiv := [...]int{5, 4, 6, 7, 3}
	//Теперь ссылается на massiv
	//Вместо переноса каждого элемента можно сделать так, тоже слайс :)
	newSlice := massiv[:]
	//Ссылка одна и та же, изменения и коснуться massiv
	newSlice[2] = 999
	fmt.Println(massiv, newSlice)

	//После этого слайс превысит len и cap, переаллоцируется в другом месте памяти, больше не будет ссылаться на "massiv"
	newSlice = append(newSlice, 228)
	//Изменения только на переаллоцированном слайсе, не перенеслись как должны были на массив
	newSlice[1] = 1337
	fmt.Println(massiv, newSlice)

	//Удалили последний элемент
	newSlice = newSlice[:len(newSlice)-1]
	fmt.Println(newSlice)

	//region #6 Exercise
	Exercise6("пупа", "лупа", "арбуzzz", "негр")
}

func Exercise6(strs ...string) {
	strArr := make([]string, len(strs))
	copy(strArr, strs)

	max := ""
	for _, v := range strArr {
		if utf8.RuneCountInString(v) > utf8.RuneCountInString(max) {
			fmt.Println(utf8.RuneCountInString(v))
			max = v
		}
	}

	fmt.Println(max)
}
