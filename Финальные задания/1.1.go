package main

import "fmt"

func main() {
	var p1, p2, p3, p4, p5 string
	var amount int

	p1, p2, p3, p4, p5 = "-", "-", "-", "-", "-"

	for {
		var str string
		var num int
		fmt.Scanln(&str, &num)
		if str == "очередь" {
			fmt.Println("1.", p1)
			fmt.Println("2.", p2)
			fmt.Println("3.", p3)
			fmt.Println("4.", p4)
			fmt.Println("5.", p5)
		} else if str == "количество" {
			fmt.Println("Осталось свободных мест:", 5 - amount)
			fmt.Println("Всего человек в очереди:", amount)
		} else if str == "конец" {
			fmt.Println("1.", p1)
			fmt.Println("2.", p2)
			fmt.Println("3.", p3)
			fmt.Println("4.", p4)
			fmt.Println("5.", p5)
			return
		} else if str == "" {
			continue
		} else if num < 1 || num > 5 {
			fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
		} else if p1 != "-" && p2 != "-" && p3 != "-" && p4 != "-" && p5 != "-" {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
		} else if num == 1 {
			if p1 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			} else {
				p1 = str
				amount++
			}
		} else if num == 2 {
			if p2 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			} else {
				p2 = str
				amount++
			}
		} else if num == 3 {
			if p3 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			} else {
				p3 = str
				amount++
			}
		} else if num == 4 {
			if p4 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			} else {
				p4 = str
				amount++
			}
		} else if num == 5 {
			if p5 != "-" {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			} else {
				p5 = str
				amount++
			}
		}
	}
}