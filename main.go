package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ExpressionResult struct { // для хранения выражения. если это арабские цифры/числа, то срез римских цифр пустой,
	// и наоборот
	Operation    string
	RomanDigits  []string
	ArabicDigits []int
}

func main() {

	expressionResult := getExpression()

	if len(expressionResult.RomanDigits) == 0 { // если это арабские цифры/числа
		fmt.Println(EvaluateDigitsExpression(expressionResult.ArabicDigits, expressionResult.Operation))
	} else { // если это римские цифры/числа
		fmt.Println(EvaluateRomansExpression(expressionResult.RomanDigits, expressionResult.Operation))
	}

}

func getExpression() ExpressionResult { // получить из строки выражение в виде ранее обозначенной структуры

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите операцию и 2 числа/цифры")

	text, _ := reader.ReadString('\n')

	splittedInput := strings.Fields(text)

	if len(splittedInput) != 3 { // если на вход подается просто строка или выражение, отличное от шаблона
		panic("Your expression is either too long or too short")
	}

	firstDigit, err1 := strconv.Atoi(splittedInput[0])
	secondDigit, err2 := strconv.Atoi(splittedInput[2])

	if reflect.TypeOf(err1) == reflect.TypeOf(err2) { // если ошибки одинакового типа (strconv.NumError | nil)
		if err1 == nil { // если при конвертации нет ошибок, это арабское число/цифра
			return ExpressionResult{Operation: splittedInput[1], ArabicDigits: []int{firstDigit, secondDigit}}
		} else { // если есть, то это римская цифра/число
			return ExpressionResult{Operation: splittedInput[1], RomanDigits: []string{splittedInput[0], splittedInput[2]}}
		}
	} else { // если ошибки разные, то и операнды разные
		panic("Operands are of different type")
	}
}
