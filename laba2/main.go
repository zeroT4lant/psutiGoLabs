package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//region #1 Exercise
	isEven(1)
	isEven(2)

	//region #2 Exercise
	Exercise2(-1)
	Exercise2(2)
	Exercise2(0)

	//region #3 Exercise
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//region #4 Exercise
	getLen("Ниггерс")
	getLen("Я в своем познании настолько преисполнился")

	//region #5 Exercise
	rect1 := Rectangle{5, 2}
	rect1.GetRectanglePloshad()

	//region #6 Exercise
	//Принимает любое количество значений
	avgNums(2, 3)
	avgNums(5, 6, 7)
}

func avgNums(num ...int) {
	sum := 0
	counter := 0
	for _, v := range num {
		counter++
		sum += v
	}
	fmt.Println(float64(sum) / float64(counter))
}

type Rectangle struct {
	osnovanie int
	visota    int
}

func (r Rectangle) GetRectanglePloshad() {
	fmt.Println((float64(r.osnovanie) * float64(r.visota)) / 2)
}

// Русские символы считаются за 2 при обычном len
func getLen(str string) {
	fmt.Println(utf8.RuneCountInString(str))
}

func isEven(num int) {
	if num%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}

func Exercise2(num int) {
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

}
