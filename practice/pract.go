package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func Clap() {
	fmt.Println("👏")
}
func Greet(name string) {
	fmt.Println("Hello,", name)
}
func Add(a int, b int) int {
	return a + b
}
func Double(number int) int {
	return number * 2
}

func main() {
	age := 10
	if age >= 18 {
		fmt.Println("Come in")
	} else {
		println("Too young")
	}
	score := 70
	if score >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	for i := 0; i < 4; i++ {
		fmt.Println("Go")
	}
	toys := [3]string{"car", "ball", "robot"}
	fmt.Println(toys)
	for i := 0; i < 3; i++ {
		fmt.Println(toys[i])
	}
	fruits := [3]string{"apple", "banana", "mango"}
	fmt.Println(fruits[1])
	fruit := []string{"banana", "apple"}
	fruit = append(fruit, "orange")
	fmt.Println(fruit)
	names1 := []string{"Mary", "John", "Peter"}
	fmt.Println(names1[0])
	scores := map[string]int{
		"Mary":  90,
		"John":  80,
		"Peter": 95,
	}
	fmt.Println(scores["Peter"])

	word := "GO"

	for _, char := range word {
		fmt.Println(char)
	}
	var err error = nil
	if err != nil {
		fmt.Println("Problem")
	} else {
		fmt.Println("Everything is okay")
	}
	SayHello()
	Clap()
	Clap()
	Greet("Adanu!")
	Greet("Gabriel!!")
	result := Add(2, 4)
	fmt.Println(result)
	answer := Double(5)
	fmt.Println(answer)

}
