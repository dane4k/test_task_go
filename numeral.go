package main

var operationsMap = map[string]func(int, int) int{ // мапа операций, чтобы без лишних if'ов быстро получить результат
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func EvaluateDigitsExpression(digits []int, operation string) int { // получить результат выражения из арабских цифр/чисел
	_, isInMap := operationsMap[operation]
	if isInMap {
		return operationsMap[operation](digits[0], digits[1])
	}
	panic("Invalid operation symbol")
}
