package main

import (
	"fmt"
	"strings"
)

func main() {
	//region #1 Exercise
	mapa := map[string]int{
		"Лена Головач":   18,
		"Артур Пирожков": 22,
		"Юрий Алышев":    228,
		"Пудж Чинчопов":  99,
	}

	//region #2 Exercise
	//Мапа под капотом сортирует по первой букве ключа, если бы были числа, то начиная с наименьшего
	fmt.Println(mapa, mapa["Юрий Алышев"])
	avgAge(mapa)

	//region #3 Exercise
	//Возвращать значение не обязательно, ссылка на исходную мапу передалась
	deleteFromMapByName(mapa, "Лена Головач")
	fmt.Println(mapa)

	//region #4 Exercise
	fmt.Println("#4 Exercise")
	//Сканит строку до пробела
	var str string
	fmt.Scan(&str)
	makeUppercase(str)

	//region #5 Exercise
	fmt.Println("#5 Exercise")
	//Введи 777 для окончания ввода чисел
	var nums []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 777 {
			break
		}

		nums = append(nums, input)
	}

	megaSum(nums)

	//region #6 Exercise
	//Введи 777 для окончания ввода чисел
	fmt.Println("#6 Exercise")
	var nums2 []int
	for {
		var input int
		fmt.Scan(&input)

		if input == 777 {
			break
		}

		nums2 = append(nums2, input)
	}
	reverseArrOfNums(nums2)
}

func reverseArrOfNums(nums []int) {
	for _, v := range nums {
		defer fmt.Println(v)
	}
}

func megaSum(nums []int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}

func makeUppercase(str string) {
	fmt.Println(strings.ToUpper(str))
}

func deleteFromMapByName(mapa map[string]int, str string) {
	fmt.Println("Before: ", mapa)
	delete(mapa, str)
	fmt.Println("After: ", mapa)
}

func avgAge(mapa map[string]int) {
	ages := 0
	counter := 0
	for _, v := range mapa {
		ages += v
		counter++
	}
	fmt.Println(float64(ages) / float64(counter))
}
