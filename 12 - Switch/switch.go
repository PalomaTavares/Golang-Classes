package main

import "fmt"

func weekdays(n int) string {
	switch n {
	case 1:
		return "Sunday - Too late"
	case 2:
		return "Monday - blue"
	case 3:
		return "Tuesday - grey"
	case 4:
		return "Wednesday - grey too"
	case 5:
		return "Thursday - I don't care"
	case 6:
		return "Friday - I'm in love"
	case 7:
		return "Saturday - Wait"
	default:
		return "Morrisey"
	}
}

func weekdays2(n int) string {
	var weekday string
	switch {
	case n == 1:
		weekday = "Sunday - Too late"
		fallthrough //cai dentro da proxima condicao
	case n == 2:
		weekday = "Monday - blue"
	case n == 3:
		weekday = "Tuesday - grey"
	case n == 4:
		weekday = "Wednesday - grey too"
	case n == 5:
		weekday = "Thursday - I don't care"
	case n == 6:
		weekday = "Friday - I'm in love"
	case n == 7:
		weekday = "Saturday - Wait"
	default:
		weekday = "Morrisey"
	}
	return weekday
}

func main() {
	fmt.Println("Switch")
	day := weekdays(6)
	day2 := weekdays2(1)
	fmt.Println(day)
	fmt.Println(day2)
}
