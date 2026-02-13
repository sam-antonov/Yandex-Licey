package main

import (
	"fmt"
	"time"
	"math"
)

func main() {
	var dateStr, name, surname, lastname string
	var n1, n2, n3 float64

	fmt.Scanln(&dateStr)
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&lastname)
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Scanln(&n3)

	beginWorkDate, _ := time.Parse("02.01.2006", dateStr)
	contractDate := beginWorkDate.AddDate(0, 0, 15)

	suma := n1 + n2 + n3

	rubles, kops := math.Modf(suma)

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n", surname, name, lastname)
	fmt.Printf("Дата подписания договора: %v. Просим вас подойти в офис в любое удобное для вас время в этот день.\n", contractDate.Format("02.01.2006"))
	fmt.Printf("Общая сумма выплат составит %d руб. %d коп.\n", int(rubles), int(math.Round(kops * 100)))
	
	fmt.Print("\nС уважением,\nГл. бух. Иванов А.Е.")
}